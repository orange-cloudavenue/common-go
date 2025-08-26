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
	"strings"
)

// Converts a string to snake_case
func ToSnake(s string) string {
	s = strings.Trim(s, " ")
	n := ""
	for i, v := range s {
		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		nextCaseIsChanged := false
		if i+1 < len(s) {
			next := s[i+1]
			if (isUpperLetter(v) && isLowerLetter(int32(next))) || (isLowerLetter(v) && isUpperLetter(int32(next))) {
				nextCaseIsChanged = true
			}
		}

		switch {
		case i > 0 && n[len(n)-1] != '_' && nextCaseIsChanged:
			if isUpperLetter(v) {
				n += "_" + string(v)
			} else if isLowerLetter(v) {
				n += string(v) + "_"
			}
		case v == ' ' || v == '-':
			n += "_"
		default:
			n += string(v)
		}
	}
	n = strings.ToLower(n)
	return n
}
