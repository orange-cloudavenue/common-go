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
	"fmt"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/orange-cloudavenue/common-go/regex"
)

func init() {
	// Register the CAV resource name generator function
	// Usage: {resource_name:edgegateway}
	gofakeit.AddFuncLookup("resource_name", gofakeit.Info{
		Category:    "cloudavenue",
		Display:     "ResourceName",
		Description: "Generate a new CAV resource name",
		Example:     "tn01e02ocb0001234spt101",
		Output:      "string",
		Params: []gofakeit.Param{
			{Field: "ress", Display: "CAV Resource Name", Type: "string", Description: "The name of the CAV resource (e.g., edgegateway, t0, etc.)", Options: func() []string {
				listOfAllowedValues := make([]string, 0)
				for _, r := range regex.ListCavResourceNames {
					listOfAllowedValues = append(listOfAllowedValues, r.Key)
				}
				return listOfAllowedValues
			}()},
		},
		Generate: func(f *gofakeit.Faker, params *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			param, err := info.GetString(params, "ress")
			if err != nil {
				return "", err
			}

			var re string
			for _, r := range regex.ListCavResourceNames {
				if r.Key == param {
					re = r.RegexString
					break
				}
			}
			if re == "" {
				return "", fmt.Errorf("unknown CAV resource name: %s", param)
			}

			// Generate a random CAV resource name based on the regex pattern
			return f.Regex(re), nil
		},
	})
}
