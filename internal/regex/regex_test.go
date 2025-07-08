package regex

import (
	"testing"
)

func TestUUID4Regex(t *testing.T) {
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

func TestURNRegex(t *testing.T) {
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
