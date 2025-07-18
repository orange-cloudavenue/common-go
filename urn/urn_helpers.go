/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package urn

import (
	"errors"
	"fmt"
)

// FindURNTypeFromString returns the URN type from a string.
func FindURNTypeFromString(value string) (URN, error) {
	if value == "" {
		return "", errors.New("value does not contain an URN type provided")
	}

	if u, ok := URNByNames[value]; ok {
		return u, nil
	}

	return "", fmt.Errorf("URN type %s does not exist in package urn", value)
}
