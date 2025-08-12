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
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/orange-cloudavenue/common-go/strcase"
)

var (

	// RequireIfNull is a validator that requires a field if another field is null.
	// It allows the field to be empty if the target field is set.
	// Param is the name or a list of names of the target fields (separated by spaces)
	// The name of the target field are in format 'GolangLike (EdgegatewayID) or paramsSpec (edgegateway_id)'.
	// Usage: `validate:"required_if_null=target_field"`
	// E.g. `validate:"required_if_null=EdgegatewayID"`
	RequireIfNull = &CustomValidator{
		Key: "required_if_null",
		Func: func(fl validator.FieldLevel) bool {
			// Field is already set, no need to validate
			if fl.Field().String() != "" {
				return true
			}

			target := fl.Param()
			if strings.Contains(target, " ") {
				// If the target is a list of fields, check if any of them is set
				for _, field := range strings.Split(target, " ") {
					if fl.Parent().FieldByName(strcase.ToPublicGoName(field)).String() != "" {
						return true
					}
				}
				return false
			}

			// Check if the parent field is null
			if fl.Parent().FieldByName(strcase.ToPublicGoName(target)).String() == "" {
				return false
			}

			return true
		},
	}

	// ExcludeIfNull is a validator that excludes a field if another field is null.
	// It allows the field to be empty if the target field is not set.
	// Param is the name or a list of names of the target fields (separated by space)
	// The name of the target field are in format 'GolangLike (EdgegatewayID) or paramsSpec (edgegateway_id)'.
	// Usage: `validate:"excluded_if_null=target_field"`
	// E.g. `validate:"excluded_if_null=EdgegatewayID"`
	ExcludeIfNull = &CustomValidator{
		Key: "excluded_if_null",
		Func: func(fl validator.FieldLevel) bool {
			// Field is already unset, no need to validate
			if fl.Field().String() == "" {
				return true
			}

			target := fl.Param()
			if strings.Contains(target, " ") {
				// If the target is a list of fields, check if any of them is set
				for _, field := range strings.Split(target, " ") {
					if fl.Parent().FieldByName(strcase.ToPublicGoName(field)).String() == "" {
						return false
					}
				}
				return true
			}

			// Check if the parent field is null
			if fl.Parent().FieldByName(strcase.ToPublicGoName(target)).String() == "" {
				return false
			}

			return true
		},
	}
)
