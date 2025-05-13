/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package utils

import (
	"fmt"
	"regexp"
)

var (
	ErrNoMatch      = fmt.Errorf("no match found")
	ErrEntryIsEmtpy = fmt.Errorf("entry is empty")
)

// GetUUIDFromHref returns the UUID from an href
// idAtEnd is true if the UUID is at the end of the href
// if href is empty, an error is returned (ErrEntryIsEmtpy)
// if no match is found, an error is returned (ErrNoMatch)
func GetUUIDFromHref(href string, idAtEnd bool) (string, error) {
	if href == "" {
		return "", ErrEntryIsEmtpy
	}

	regex := `:\/\/.+([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})`

	if idAtEnd {
		regex += `$`
	} else {
		regex += `.*$`
	}

	reGetID := regexp.MustCompile(regex)
	matchList := reGetID.FindAllStringSubmatch(href, -1)

	if len(matchList) == 0 {
		return "", ErrNoMatch
	}

	return matchList[0][1], nil
}
