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

	"github.com/go-playground/validator/v10"
)

// FormatValidator is a custom validator that checks if a string matches a given format (e.g. email)
var FormatValidator = &CustomValidator{
	Key: "format",
	Func: func(fl validator.FieldLevel) bool {
		if fl.Field().Kind() != reflect.String {
			return false
		}
		switch fl.Param() {
		case "email":
			v := validator.New()
			return v.Var(fl.Field().String(), "email") == nil
		default:
			return false
		}
	},
}
