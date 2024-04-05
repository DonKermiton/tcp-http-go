package tcp_http_go

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expected      map[string]string
		expectedError error
	}{
		{
			name: "Valid headers with spaces",
			input: "Header1:  value1  \r\n" +
				"Header2: value2\r\n" +
				"\r\n",
			expected: map[string]string{
				"Header1": "value1",
				"Header2": "value2",
			},
		},
		{
			name:     "Empty input",
			input:    "",
			expected: map[string]string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h := &headers{headers: make(map[string]string)}
			reader := bytes.NewReader([]byte(tc.input))

			err := h.decode(reader)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}

			if !reflect.DeepEqual(h.headers, tc.expected) {
				t.Errorf("Expected headers: %v, got: %v", tc.expected, h.headers)
			}
		})
	}
}
