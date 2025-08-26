/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package regex

import (
	"regexp"
	"sync"
)

const (
	UUID4RegexString        = `[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}`
	URNWithUUID4RegexString = `(?m)urn:[A-Za-z0-9][A-Za-z0-9-]{0,31}:([A-Za-z0-9()+,\-.:=@;$_!*']|%[0-9A-Fa-f]{2})+` + UUID4RegexString

	// * Cloudavenue
	// Resource Names Regex declared in internal/regex/resource_name.go

	// * Cases
	PascalCaseRegexString = `^[A-Z](([a-z0-9]+[A-Z]?)*)$`
	CamelCaseRegexString  = `^[a-z](([a-z0-9]+[A-Z]?)*)$`
	SnakeCaseRegexString  = `^[a-z0-9]+(_[a-z0-9]+)*$`
	KebabCaseRegexString  = `^[a-z0-9]+(-[a-z0-9]+)*$`
	UpperCaseRegexString  = `^[A-Z]+(_[A-Z]+)*$`
)

func lazyRegexCompile(str string) func() *regexp.Regexp {
	var regex *regexp.Regexp
	var once sync.Once
	return func() *regexp.Regexp {
		once.Do(func() {
			regex = regexp.MustCompile(str)
		})
		return regex
	}
}

var (
	UUID4Regex        = lazyRegexCompile(UUID4RegexString)
	URNWithUUID4Regex = lazyRegexCompile(URNWithUUID4RegexString)

	// * Cloudavenue
	T0NameRegex          = lazyRegexCompile(T0NameRegexString)
	EdgeGatewayNameRegex = lazyRegexCompile(EdgeGatewayNameRegexString)

	// * Cases
	PascalCaseRegex = lazyRegexCompile(PascalCaseRegexString)
	CamelCaseRegex  = lazyRegexCompile(CamelCaseRegexString)
	SnakeCaseRegex  = lazyRegexCompile(SnakeCaseRegexString)
	KebabCaseRegex  = lazyRegexCompile(KebabCaseRegexString)
	UpperCaseRegex  = lazyRegexCompile(UpperCaseRegexString)
)
