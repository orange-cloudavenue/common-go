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

	"github.com/orange-cloudavenue/common-go/urn"
)

func init() {
	gofakeit.AddFuncLookup("urn", gofakeit.Info{
		Category:    "custom",
		Display:     "urn",
		Description: "Generate a new URN (with custom type)",
		Example:     "urn:vcloud:gateway:12345678-1234-4234-1234-123456789012",
		Output:      "string",
		Params: []gofakeit.Param{
			{Field: "urnType", Type: "string", Description: "The type of the URN, e.g., 'edgegateway', 'vdc', etc."},
		},
		Generate: func(f *gofakeit.Faker, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			urnType, _ := info.GetString(m, "urnType")

			newURN, err := urn.FindURNTypeFromString(urnType)
			if err != nil {
				return nil, err
			}

			return urn.Normalize(newURN, f.UUID()).String(), nil
		},
	})
}
