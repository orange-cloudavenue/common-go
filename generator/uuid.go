/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package generator

import "github.com/brianvoe/gofakeit/v7"

func init() {
	gofakeit.AddFuncLookup("uuid", gofakeit.Info{
		Category: "custom",
		Display:  "uuid",
		Generate: func(f *gofakeit.Faker, _ *gofakeit.MapParams, _ *gofakeit.Info) (any, error) {
			return f.UUID(), nil
		},
	})
}
