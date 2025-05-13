/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package print

import (
	"fmt"
	"testing"
)

func TestFieldFormat(t *testing.T) {
	testCases := []struct {
		value    interface{}
		expected string
	}{
		{
			value:    []int{1, 2, 3},
			expected: "1,2,3",
		},
		{
			value:    map[string]int{"a": 1, "b": 2},
			expected: "a:1,b:2",
		},
		{
			value: struct {
				Name  string
				Age   int
				Email string
			}{
				Name:  "John",
				Age:   30,
				Email: "john@example.com",
			},
			expected: "Name:John,Age:30,Email:john@example.com",
		},
		{
			value:    true,
			expected: "Enabled",
		},
		{
			value:    false,
			expected: "Disabled",
		},
	}

	for _, tc := range testCases {
		actual := fieldFormat(tc.value)
		if fmt.Sprintf("%v", actual) != tc.expected {
			t.Errorf("Expected FieldFormat(%v) to return %s, but got %v", tc.value, tc.expected, actual)
		}
	}
}

func TestFormatEntry(t *testing.T) {
	testCases := []struct {
		value    interface{}
		expected string
	}{
		{
			value:    "Hello",
			expected: String.String() + Tab.String(),
		},
		{
			value:    42,
			expected: Int.String() + Tab.String(),
		},
		{
			value:    3.14,
			expected: Float.String() + Tab.String(),
		},
		{
			value:    true,
			expected: Bool.String() + Tab.String(),
		},
		{
			value:    struct{}{},
			expected: Default.String() + Tab.String(),
		},
	}

	for _, tc := range testCases {
		actual := formatEntry(tc.value)
		if actual != tc.expected {
			t.Errorf("Expected formatEntry(%v) to return %s, but got %s", tc.value, tc.expected, actual)
		}
	}
}
func TestFormat(t *testing.T) {
	testCases := []struct {
		fields   []interface{}
		expected string
	}{
		{
			fields:   []interface{}{"Hello", 42, 3.14, true, struct{}{}},
			expected: String.String() + Tab.String() + Int.String() + Tab.String() + Float.String() + Tab.String() + Bool.String() + Tab.String() + Default.String() + Tab.String() + NewLine.String(),
		},
		{
			fields:   []interface{}{1, 2, 3, "a", "b", "c"},
			expected: Int.String() + Tab.String() + Int.String() + Tab.String() + Int.String() + Tab.String() + String.String() + Tab.String() + String.String() + Tab.String() + String.String() + Tab.String() + NewLine.String(),
		},
		{
			fields:   []interface{}{true, false, true, false},
			expected: Bool.String() + Tab.String() + Bool.String() + Tab.String() + Bool.String() + Tab.String() + Bool.String() + Tab.String() + NewLine.String(),
		},
	}

	for _, tc := range testCases {
		actual := format(tc.fields...)
		if actual != tc.expected {
			t.Errorf("Expected format(%v) to return %s, but got %s", tc.fields, tc.expected, actual)
		}
	}
}
