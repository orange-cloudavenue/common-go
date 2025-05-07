/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package main

import "github.com/orange-cloudavenue/common-go/print"

func main() {
	p := print.New()
	p.SetHeader("String", "Int", "Bool", "Slice", "Map", "Struct", "Array")
	p.AddFields("I'm a string", 1, true, []string{"a", "b", "c"}, map[string]string{"a": "1", "b": "2", "c": "3"}, struct{ key1, key2, key3 string }{"a", "b", "c"}, [3]string{"a", "b", "c"})
	p.PrintTable()
}
