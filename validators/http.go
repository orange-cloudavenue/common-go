/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package validators

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// HTTPStatusCode is a custom validator that checks if a string is a valid HTTP status code.
var HTTPStatusCode = &CustomValidator{
	Key: "http_status_code",
	Func: func(fl validator.FieldLevel) bool {
		var (
			i   int
			err error
		)

		if fl.Field().Type().Kind() == reflect.String {
			// convert string to int
			i, err = strconv.Atoi(fl.Field().String())
			if err != nil {
				return false
			}
		} else {
			i = int(fl.Field().Int())
		}

		// check if the integer is a valid HTTP status code
		if i < 100 || i > 599 {
			return false
		}

		return true
	},
}

// HTTPStatusCodeRange is a custom validator that checks if a string is a valid HTTP status code range.
var HTTPStatusCodeRange = &CustomValidator{
	Key: "http_status_code_range",
	Func: func(fl validator.FieldLevel) bool {
		// format of the string is "100-599"
		// split the string into two parts
		p := strings.Split(fl.Field().String(), "-")

		if len(p) != 2 {
			return false
		}

		// convert the strings to integers
		start, err := strconv.Atoi(p[0])
		if err != nil {
			return false
		}

		end, err := strconv.Atoi(p[1])
		if err != nil {
			return false
		}

		// check if the integers are valid HTTP status codes
		if start < 100 || start > 599 {
			return false
		}

		if end < 100 || end > 599 {
			return false
		}

		// check if the start is less than the end
		if start >= end {
			return false
		}

		return true
	},
}
