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

	"github.com/orange-cloudavenue/common-go/urn"
)

func TestGenerator_URN(t *testing.T) {
	type CStruct struct {
		ID string `fake:"{urn:vm}"`
	}

	var st CStruct

	err := Struct(&st)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	if !urn.IsVM(st.ID) {
		t.Fatalf("Expected URN to be of type VM, got %s", st.ID)
	}
}

func TestGenerator_URN_NoURNType(t *testing.T) {
	type CStruct struct {
		ID string `fake:"{urn}"`
	}

	var st CStruct

	err := Struct(&st)
	if err == nil {
		t.Fatal("Expected error when no URN type is provided, but got none")
	}
}
