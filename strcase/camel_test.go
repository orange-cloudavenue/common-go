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

func TestToCamel(t *testing.T) {
	cases := [][]string{
		{"foo-bar", "fooBar"},
		{"TestCase", "testCase"},
		{"", ""},
		{"AnyKind of_string", "anyKindOfString"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToCamel(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}
