package regex

import (
	"testing"
)

func TestRegex_UUID4(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError bool
	}{
		{
			name:          "Valid UUID4",
			input:         "d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expected:      "d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expectedError: false,
		},
		{
			name:          "Invalid UUID4",
			input:         "d3c42a20-96b9-4452-91dd-f71b71dfe31Z",
			expected:      "",
			expectedError: true,
		},
		{
			name:          "Empty string",
			input:         "",
			expected:      "",
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			matches := UUID4Regex().FindAllString(test.input, -1)
			if len(matches) == 0 && !test.expectedError {
				t.Errorf("Expected to find a match for %s, but found none", test.input)
			}
			if len(matches) > 1 && !test.expectedError {
				t.Errorf("Expected only one match for %s, but found %d", test.input, len(matches))
			}
			if len(matches) == 1 && matches[0] != test.expected {
				t.Errorf("Expected match %s, got %s", test.expected, matches[0])
			}
		})
	}
}

func TestRegex_URN(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError bool
	}{
		{
			name:          "Valid URN with UUID4",
			input:         "urn:example:service:d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expected:      "urn:example:service:d3c42a20-96b9-4452-91dd-f71b71dfe314",
			expectedError: false,
		},
		{
			name:          "Invalid URN without UUID4",
			input:         "urn:example:service:not-a-valid-uuid",
			expected:      "",
			expectedError: true,
		},
		{
			name:          "Empty string",
			input:         "",
			expected:      "",
			expectedError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			matches := URNWithUUID4Regex().FindAllString(test.input, -1)
			if len(matches) == 0 && !test.expectedError {
				t.Errorf("Expected to find a match for %s, but found none", test.input)
			}
			if len(matches) > 1 && !test.expectedError {
				t.Errorf("Expected only one match for %s, but found %d", test.input, len(matches))
			}
			if len(matches) == 1 && matches[0] != test.expected {
				t.Errorf("Expected match %s, got %s", test.expected, matches[0])
			}
		})
	}
}

func TestRegex_T0Name(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError bool
	}{
		{
			name:          "Valid T0 Name",
			input:         "prvrf01eocb0001234allsp01",
			expected:      "prvrf01eocb0001234allsp01",
			expectedError: false,
		},
		{
			name:          "Invalid T0 Name",
			input:         "invalid-name",
			expected:      "",
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			matches := T0NameRegex().FindAllString(test.input, -1)
			if len(matches) == 0 && !test.expectedError {
				t.Errorf("Expected to find a match for %s, but found none", test.input)
			}
			if len(matches) > 1 && !test.expectedError {
				t.Errorf("Expected only one match for %s, but found %d", test.input, len(matches))
			}
			if len(matches) == 1 && matches[0] != test.expected {
				t.Errorf("Expected match %s, got %s", test.expected, matches[0])
			}
		})
	}
}

func TestRegex_EdgeGatewayName(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError bool
	}{
		{
			name:          "Valid Edge Gateway Name",
			input:         "tn01e02ocb0001234spt101",
			expected:      "tn01e02ocb0001234spt101",
			expectedError: false,
		},
		{
			name:          "Invalid Edge Gateway Name",
			input:         "invalid-name",
			expected:      "",
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			matches := EdgeGatewayNameRegex().FindAllString(test.input, -1)
			if len(matches) == 0 && !test.expectedError {
				t.Errorf("Expected to find a match for %s, but found none", test.input)
			}
			if len(matches) > 1 && !test.expectedError {
				t.Errorf("Expected only one match for %s, but found %d", test.input, len(matches))
			}
			if len(matches) == 1 && matches[0] != test.expected {
				t.Errorf("Expected match %s, got %s", test.expected, matches[0])
			}
		})
	}
}

// * ----
func TestLazyRegexCompile(t *testing.T) {
	callCount := 0
	fakeCompile := func(str string) func() *string {
		var compiled *string
		var onceCalled bool
		return func() *string {
			if !onceCalled {
				onceCalled = true
				callCount++
				compiled = &str
			}
			return compiled
		}
	}
	getter := fakeCompile("abc")
	_ = getter()
	_ = getter()
	if callCount != 1 {
		t.Errorf("Expected compile to be called once, got %d", callCount)
	}
}
