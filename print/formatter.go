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
	"reflect"
	"strings"
)

type (
	fmtFormat string
)

const (
	// String fmtFormat
	String fmtFormat = "%s"
	// Int fmtFormat
	Int fmtFormat = "%d"
	// Float fmtFormat
	Float fmtFormat = "%f"
	// Bool fmtFormat
	Bool fmtFormat = "%s" // Use %s for bool values instead of %t because AddFields() converts bool to string
	// Default fmtFormat
	Default fmtFormat = "%v"

	// Tab fmtFormat
	Tab fmtFormat = "\t"
	// NewLine fmtFormat
	NewLine fmtFormat = "\n"
)

// String
func (f fmtFormat) String() string {
	return string(f)
}

// formatEntry is a function that formats the given value based on its type.
// It returns the formatted string representation of the value.
func formatEntry(value any) (fs string) {
	switch value.(type) {
	case string:
		fs += String.String()
	case int, int8, int16, int32, int64:
		fs += Int.String()
	case float32, float64:
		fs += Float.String()
	case bool:
		fs += Bool.String()
	default:
		fs += Default.String()
	}

	return fs + Tab.String()
}

// format is a function that takes in variadic arguments of any type and returns a formatted string.
// It concatenates the formatted string representation of each argument using the formatEntry function,
// and appends a newline character at the end.
func format(fields ...any) (fs string) {
	fs = ""
	for _, field := range fields {
		fs += formatEntry(field)
	}
	fs += NewLine.String()
	return fs
}

// fieldFormat
// fieldFormat is a function that formats a value based on its type.
// It takes any value as input and returns a formatted string.
// The formatting rules are as follows:
// - For slices and arrays, it joins the elements with a comma.
// - For maps, it formats the key-value pairs as "key:value" and joins them with a comma.
// - For structs, it formats the field name and value pairs as "fieldName:fieldValue" and joins them with a comma.
// - For booleans, it returns "Enabled" if the value is true, and "Disabled" if the value is false.
// For all other types, it returns the input value as is.
func fieldFormat(value any) any {
	x := reflect.ValueOf(value)
	switch x.Kind() {
	case reflect.Slice, reflect.Array:
		var s []string
		for i := 0; i < x.Len(); i++ {
			s = append(s, fmt.Sprintf("%v", x.Index(i)))
		}
		return strings.Join(s, ",")
	case reflect.Map:
		// format map to "a:1,b:2"
		var s []string
		for _, k := range x.MapKeys() {
			s = append(s, fmt.Sprintf("%v:%v", k, x.MapIndex(k)))
		}
		return strings.Join(s, ",")
	case reflect.Struct:
		// format struct to "fieldName:fieldValue,fieldName:fieldValue"
		var s []string
		for i := 0; i < x.NumField(); i++ {
			s = append(s, fmt.Sprintf("%v:%v", x.Type().Field(i).Name, x.Field(i)))
		}
		return strings.Join(s, ",")
	case reflect.Bool:
		if x.Bool() {
			return "Enabled"
		}
		return "Disabled"
	}

	return value
}
