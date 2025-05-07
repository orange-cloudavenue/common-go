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
	"testing"
)

func TestGetUUIDFromHref(t *testing.T) {
	tests := []struct {
		name        string
		href        string
		idAtEnd     bool
		expectedID  string
		expectedErr error
	}{
		{
			name:        "Valid href with ID at the end",
			href:        "https://example.com/resource/123e4567-e89b-12d3-a456-426614174000",
			idAtEnd:     true,
			expectedID:  "123e4567-e89b-12d3-a456-426614174000",
			expectedErr: nil,
		},
		{
			name:        "Valid href with ID not at the end",
			href:        "https://example.com/resource/123e4567-e89b-12d3-a456-426614174000/something",
			idAtEnd:     false,
			expectedID:  "123e4567-e89b-12d3-a456-426614174000",
			expectedErr: nil,
		},
		{
			name:        "Invalid href",
			href:        "https://example.com/resource/invalid",
			idAtEnd:     true,
			expectedID:  "",
			expectedErr: ErrNoMatch,
		},
		{
			name:        "Empty href",
			href:        "",
			idAtEnd:     true,
			expectedID:  "",
			expectedErr: ErrEntryIsEmtpy,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, err := GetUUIDFromHref(test.href, test.idAtEnd)

			if id != test.expectedID {
				t.Errorf("Expected ID: %s, but got: %s", test.expectedID, id)
			}

			if err != test.expectedErr {
				t.Errorf("Expected error: %v, but got: %v", test.expectedErr, err)
			}
		})
	}
}
