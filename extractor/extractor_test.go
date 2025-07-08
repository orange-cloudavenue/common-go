package extractor

import (
	"testing"
)

func TestExtractUUID_Success(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError bool
	}{
		{
			name:          "Valid UUID in prefix",
			input:         "prefix d3c42a20-96b9-4452-91dd-f71b71dfe314 suffix",
			expected:      "d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expectedError: false,
		},
		{
			name:          "Valid UUID in URL at the end",
			input:         "https://console1.cloudavenue.orange-business.com/api/admin/edgeGateway/d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expected:      "d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expectedError: false,
		},
		{
			name:          "Valid UUID in URL in the middle",
			input:         "https://console1.cloudavenue.orange-business.com/api/admin/edgeGateway/d3c42a20-96b9-4452-91dd-f71b71dfe314/in/the/middle",
			expected:      "d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expectedError: false,
		},
		{
			name:          "Multiple UUIDs in input",
			input:         "prefix d3c42a20-96b9-4452-91dd-f71b71dfe314 and another d3c42a20-96b9-4452-91dd-f71b71dfe315",
			expected:      "",
			expectedError: true,
		},
		{
			name:          "No UUID in input",
			input:         "this string has no uuid",
			expected:      "",
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ExtractUUID(test.input)
			if (err != nil) != test.expectedError {
				t.Fatalf("expected error: %v, got: %v", test.expectedError, err)
			}
			if result != test.expected {
				t.Fatalf("expected: %s, got: %s", test.expected, result)
			}
		})
	}
}

func TestExtractURN(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError bool
	}{
		{
			name:          "Valid URN with UUID4",
			input:         "urn:example:resource:d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expected:      "urn:example:resource:d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expectedError: false,
		},
		{
			name:          "URN with UUID4 in the middle of text",
			input:         "prefix urn:example:resource:d3c42a20-96b9-4452-91dd-f71b71dfe314 suffix",
			expected:      "urn:example:resource:d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expectedError: false,
		},
		{
			name:          "Multiple URNs with UUID4",
			input:         "urn:example:resource:d3c42a20-96b9-4452-91dd-f71b71dfe314 and urn:example:resource:d3c42a20-96b9-4452-91dd-f71b71dfe315",
			expected:      "",
			expectedError: true,
		},
		{
			name:          "Valid URN with UUID4 in URL",
			input:         "https://console1.cloudavenue.orange-business.com/api/admin/edgeGateway/urn:example:resource:d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expected:      "urn:example:resource:d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expectedError: false,
		},
		{
			name:          "No URN with UUID4",
			input:         "this string has no urn",
			expected:      "",
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ExtractURN(test.input)
			if (err != nil) != test.expectedError {
				t.Fatalf("expected error: %v, got: %v", test.expectedError, err)
			}
			if result != test.expected {
				t.Fatalf("expected: %s, got: %s", test.expected, result)
			}
		})
	}
}
