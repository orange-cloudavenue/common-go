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
	"testing"
)

func TestFindURNTypeFromString(t *testing.T) {
	// Setup a fake URNByNames for testing
	URNByNames = map[string]URN{
		"test-type": "test-type",
	}

	tests := []struct {
		name      string
		input     string
		want      URN
		wantErr   bool
		errString string
	}{
		{
			name:      "empty value",
			input:     "",
			want:      "",
			wantErr:   true,
			errString: "value does not contain an URN type provided",
		},
		{
			name:    "existing type",
			input:   "test-type",
			want:    "test-type",
			wantErr: false,
		},
		{
			name:      "non-existing type",
			input:     "not-exist",
			want:      "",
			wantErr:   true,
			errString: "URN type not-exist does not exist in package urn",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindURNTypeFromString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindURNTypeFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindURNTypeFromString() = %v, want %v", got, tt.want)
			}
			if tt.wantErr && err != nil && tt.errString != "" {
				if !errors.Is(err, errors.New(tt.errString)) && err.Error() != tt.errString {
					t.Errorf("FindURNTypeFromString() error = %v, want %v", err, tt.errString)
				}
			}
		})
	}
}
