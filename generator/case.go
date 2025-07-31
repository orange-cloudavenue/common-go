/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package generator

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/orange-cloudavenue/common-go/internal/regex"
)

func init() {
	// Register the camelCase generator function
	gofakeit.AddFuncLookup("camelCase", gofakeit.Info{
		Category:    "cloudavenue",
		Display:     "CamelCase",
		Description: "Generate a new camelCase string",
		Example:     "camelCaseExample",
		Output:      "string",
		Generate: func(f *gofakeit.Faker, _ *gofakeit.MapParams, _ *gofakeit.Info) (any, error) {
			// Generate a random camelCase string
			return f.Regex(regex.CamelCaseRegexString), nil
		},
	})

	// Register the PascalCase generator function
	gofakeit.AddFuncLookup("PascalCase", gofakeit.Info{
		Category:    "cloudavenue",
		Display:     "PascalCase",
		Description: "Generate a new PascalCase string",
		Example:     "PascalCaseExample",
		Output:      "string",
		Generate: func(f *gofakeit.Faker, _ *gofakeit.MapParams, _ *gofakeit.Info) (any, error) {
			// Generate a random PascalCase string
			return f.Regex(regex.PascalCaseRegexString), nil
		},
	})

	// Register the snake_case generator function
	gofakeit.AddFuncLookup("snake_case", gofakeit.Info{
		Category:    "cloudavenue",
		Display:     "SnakeCase",
		Description: "Generate a new snake_case string",
		Example:     "snake_case_example",
		Output:      "string",
		Generate: func(f *gofakeit.Faker, _ *gofakeit.MapParams, _ *gofakeit.Info) (any, error) {
			// Generate a random snake_case string
			return f.Regex(regex.SnakeCaseRegexString), nil
		},
	})

	// Register the UPPER_CASE generator function
	gofakeit.AddFuncLookup("UPPER_CASE", gofakeit.Info{
		Category:    "cloudavenue",
		Display:     "UpperCase",
		Description: "Generate a new UPPER_CASE string",
		Example:     "UPPER_CASE_EXAMPLE",
		Output:      "string",
		Generate: func(f *gofakeit.Faker, _ *gofakeit.MapParams, _ *gofakeit.Info) (any, error) {
			// Generate a random UPPER_CASE string
			return f.Regex(regex.UpperCaseRegexString), nil
		},
	})

	// Register the kebab-case generator function
	gofakeit.AddFuncLookup("kebab-case", gofakeit.Info{
		Category:    "cloudavenue",
		Display:     "KebabCase",
		Description: "Generate a new kebab-case string",
		Example:     "kebab-case-example",
		Output:      "string",
		Generate: func(f *gofakeit.Faker, _ *gofakeit.MapParams, _ *gofakeit.Info) (any, error) {
			// Generate a random kebab-case string
			return f.Regex(regex.KebabCaseRegexString), nil
		},
	})
}
