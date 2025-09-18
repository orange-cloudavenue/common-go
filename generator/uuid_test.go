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
	"testing"

	"github.com/orange-cloudavenue/common-go/regex"
)

func TestGenerator_UUID(t *testing.T) {
	type CStruct struct {
		Uuid string `fake:"{uuid}"`
	}

	var st CStruct

	err := Struct(&st)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	if !regex.UUID4Regex().MatchString(st.Uuid) {
		t.Fatalf("Expected to have a valid UUID4, got %s", st.Uuid)
	}
}
