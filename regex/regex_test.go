/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package regex

import (
	"regexp"
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

func TestRegex_OrganizationName(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError bool
	}{
		{
			name:          "Valid Organization Name",
			input:         "cav01ev01ocb0001234",
			expected:      "cav01ev01ocb0001234",
			expectedError: false,
		},
		{
			name:          "Invalid Organization Name",
			input:         "invalid-name",
			expected:      "",
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			matches := OrganizationNameRegex().FindAllString(test.input, -1)
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

// TestRegex_Cases tests the regex patterns for different naming conventions.
func TestRegex_Cases(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError bool
		regex         func() *regexp.Regexp
	}{
		{
			name:          "Valid PascalCase",
			input:         "PascalCaseExample",
			expected:      "PascalCaseExample",
			expectedError: false,
			regex:         PascalCaseRegex,
		},
		{
			name:          "Invalid PascalCase",
			input:         "pascalcaseexample",
			expected:      "",
			expectedError: true,
			regex:         PascalCaseRegex,
		},
		{
			name:          "Valid CamelCase",
			input:         "camelCaseExample",
			expected:      "camelCaseExample",
			expectedError: false,
			regex:         CamelCaseRegex,
		},
		{
			name:          "Invalid CamelCase",
			input:         "CamelCaseExample",
			expected:      "",
			expectedError: true,
			regex:         CamelCaseRegex,
		},
		{
			name:          "Valid Snake_Case",
			input:         "snake_case_example",
			expected:      "snake_case_example",
			expectedError: false,
			regex:         SnakeCaseRegex,
		},
		{
			name:          "Invalid Snake_Case",
			input:         "SnakeCaseExample",
			expected:      "",
			expectedError: true,
			regex:         SnakeCaseRegex,
		},
		{
			name:          "Valid Kebab-Case",
			input:         "kebab-case-example",
			expected:      "kebab-case-example",
			expectedError: false,
			regex:         KebabCaseRegex,
		},
		{
			name:          "Invalid Kebab-Case",
			input:         "KebabCaseExample",
			expected:      "",
			expectedError: true,
			regex:         KebabCaseRegex,
		},
		{
			name:          "Valid UPPER_CASE",
			input:         "UPPER_CASE_EXAMPLE",
			expected:      "UPPER_CASE_EXAMPLE",
			expectedError: false,
			regex:         UpperCaseRegex,
		},
		{
			name:          "Invalid UPPER_CASE",
			input:         "UpperCaseExample",
			expected:      "",
			expectedError: true,
			regex:         UpperCaseRegex,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			matches := test.regex().FindAllString(test.input, -1)
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
