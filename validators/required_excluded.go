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
	"github.com/go-playground/validator/v10"
)

var (

	// RequireIfNull is a validator that requires a field if another field is null.
	RequireIfNull = &CustomValidator{
		Key: "required_if_null",
		Func: func(fl validator.FieldLevel) bool {
			// Field is already set, no need to validate
			if fl.Field().String() != "" {
				return true
			}

			// Check if the parent field is null
			if fl.Parent().FieldByName(fl.Param()).String() == "" {
				return false
			}

			return true
		},
	}

	// ExcludeIfNull is a validator that excludes a field if another field is null.
	ExcludeIfNull = &CustomValidator{
		Key: "excluded_if_null",
		Func: func(fl validator.FieldLevel) bool {
			// Field is already unset, no need to validate
			if fl.Field().String() == "" {
				return true
			}

			// Check if the parent field is null
			if fl.Parent().FieldByName(fl.Param()).String() == "" {
				return false
			}

			return true
		},
	}

	//
)
