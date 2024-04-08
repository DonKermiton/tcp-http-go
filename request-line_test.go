package tcp_http_go

import (
	"fmt"
	"testing"
)

func TestIsMethodCorrect(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Correct method", "GET", true},
		{"Invalid method", "Test", false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := isMethodCorrect(testCase.input)

			if result != testCase.expected {
				t.Errorf("Expected value: %v, got: %v", testCase.expected, result)
			}
		})
	}

}

func TestIsHttpVersionCorrect(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid http protocol http/1.0", "HTTP/1.0", true},
		{"valid http protocol http/1.1", "HTTP/1.1", true},
		{"invalid http protocol http/0.9", "HTTP/0.9", false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := isHttpVersionCorrect(testCase.input)

			if result != testCase.expected {
				t.Errorf("Expected value; %v, got: %v", result, testCase.expected)
			}
		})
	}
}

func TestSplitRequestHeaderAndValidate(t *testing.T) {
	testCases := []struct {
		name       string
		rawRequest string
		wantError  bool
	}{
		{"Valid request 2 components", "GET /hello", false},
		{"Valid request 3 components", "GET /hello HTTP/1.1", false},
		{"Invalid - too few components", "GET", true},
		{"Invalid - too many components", "GET /hello HTTP/1.1 extra", true},
		{"Invalid - empty request", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, gotErr := splitRequestHeaderAndValidate(tc.rawRequest)

			if tc.wantError && gotErr == nil {
				t.Errorf("Expected an error, but got none")
			}

			if !tc.wantError && gotErr != nil {
				t.Errorf("Unexpected error: %v", gotErr)
			}

		})
	}
}

func TestRequestLineDecode(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError error
	}{
		{name: "GET request HTTP 0.9", input: "GET /echo", expectedError: nil},
		{name: "GET request HTTP 1.1", input: "GET /echo HTTP/1.1", expectedError: nil},
		{name: "POST request HTTP 1.0", input: "POST /api/data HTTP/1.0", expectedError: nil},

		{name: "Invalid method", input: "WRONGMETHOD /path HTTP/1.1", expectedError: fmt.Errorf("invalid method")},
		{name: "Invalid HTTP version", input: "GET /path HTTP/1.2", expectedError: fmt.Errorf("invalid http protocol version")},
		{name: "Too many components", input: "GET /path HTTP/1.1 EXTRA", expectedError: fmt.Errorf("invalid number of components in request line")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rl := newRequestLine(tc.input)
			err := rl.decode()

			if err != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("expected: %v, got: %v", tc.expectedError, err.Error())
			}

		})
	}
}
