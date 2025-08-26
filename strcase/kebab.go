/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package strcase

import "strings"

// Converts a string to kebab-case
func ToKebab(s string) string {
	return strings.ReplaceAll(ToSnake(s), "_", "-")
}

// Converts a string to kebab-case
func ToSpace(s string) string {
	return strings.ReplaceAll(ToSnake(s), "_", " ")
}
