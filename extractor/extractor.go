/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package extractor

import (
	"fmt"

	"github.com/orange-cloudavenue/common-go/internal/regex"
)

// ExtractUUID extract the UUID (version 4) found in the input string using a regular expression.
// Returns the UUID as a string if found, otherwise returns an error.
func ExtractUUID(input string) (string, error) {
	matches := regex.UUID4Regex().FindAllString(input, -1)
	return helperMatch("ExtractUUID", matches, 1)
}

// ExtractURN extract the URN (Uniform Resource Name) with a UUID4 found in the input string using a regular expression.
// The URN must follow the format: urn:<namespace>:<name>:<uuid4>.
// The UUID4 is expected to be at the end of the URN.
// Returns the URN as a string if found, otherwise returns an error.
func ExtractURN(input string) (string, error) {
	matches := regex.URNWithUUID4Regex().FindAllString(input, -1)
	return helperMatch("ExtractURN", matches, 1)
}

// helperMatch checks the number of matches found against the expected number.
// Returns the first match if the count is as expected, otherwise returns an error with details.
func helperMatch(operation string, matches []string, expectedMatch int) (string, error) {
	if len(matches) == 0 {
		return "", fmt.Errorf("[%s] no matches found in input", operation)
	}

	if len(matches) != expectedMatch {
		return "", fmt.Errorf("[%s] expected %d matches, got %d: %v", operation, expectedMatch, len(matches), matches)
	}
	return matches[0], nil
}
