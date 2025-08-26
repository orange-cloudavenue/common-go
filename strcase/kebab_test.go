/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package strcase

import (
	"testing"
)

func TestToKebab(t *testing.T) {
	cases := [][]string{
		{"testCase", "test-case"},
		{"TestCase", "test-case"},
		{"Test Case", "test-case"},
		{" Test Case", "test-case"},
		{"Test Case ", "test-case"},
		{" Test Case ", "test-case"},
		{"test", "test"},
		{"test_case", "test-case"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "many-many-words"},
		{"manyManyWords", "many-many-words"},
		{"AnyKind of_string", "any-kind-of-string"},
		{"numbers2and55with000", "numbers2and55with000"},
		{"JSONData", "json-data"},
		{"userID", "user-id"},
		{"AAAbbb", "aa-abbb"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToKebab(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}
