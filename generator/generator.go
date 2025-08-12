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
)

// Reimporting gofakeit functions to the package level for access with generator package
var (

	// Struct fills in exported fields of a struct with random data based on
	// the value of `fake` tag of exported fields or with the result of a call
	// to the Fake() method if the field type implements `Fakeable`.
	// Use `fake:"skip"` to explicitly skip an element.
	// All built-in types are supported, with templating support for string types.
	Struct = gofakeit.Struct

	// Template generates an document based on the the supplied template
	Template = gofakeit.Template

	// Slice fills built-in types and exported fields of a struct with random data.
	Slice = gofakeit.Slice

	// Map will generate a random set of map data
	Map = gofakeit.Map

	// Generate fake information from given string.
	// Replaceable values should be within {}
	//
	// Functions
	// Ex: {firstname} - billy
	// Ex: {sentence:3} - Record river mind.
	// Ex: {number:1,10} - 4
	// Ex: {uuid} - 590c1440-9888-45b0-bd51-a817ee07c3f2
	//
	// Letters/Numbers
	// Ex: ### - 481 - random numbers
	// Ex: ??? - fda - random letters
	//
	// For a complete list of runnable functions use FuncsLookup
	Generate = gofakeit.Generate

	// Regex will generate a string based upon a RE2 syntax
	Regex = gofakeit.Regex
)

func MustTemplate(template string, co *gofakeit.TemplateOptions) string {
	result, err := gofakeit.Template(template, co)
	if err != nil {
		panic(err)
	}
	return result
}

func MustGenerate(template string) string {
	result, err := gofakeit.Generate(template)
	if err != nil {
		panic(err)
	}
	return result
}

func MustStruct(v any) {
	if err := gofakeit.Struct(v); err != nil {
		panic(err)
	}
}
