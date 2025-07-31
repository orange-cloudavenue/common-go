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

	"github.com/orange-cloudavenue/common-go/internal/regex"
)

func TestGenerator_CAVResourceName(t *testing.T) {
	type CStruct struct {
		Name string `fake:"{resource_name:edgegateway}"`
	}

	var st CStruct

	err := Struct(&st)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	if !regex.EdgeGatewayNameRegex().MatchString(st.Name) {
		t.Fatalf("Expected CAV resource name to match regex, got %s", st.Name)
	}
}

func TestGenerator_CAVResourceName_Invalid(t *testing.T) {
	type CStruct struct {
		Name string `fake:"{resource_name:invalid_resource}"`
	}

	var st CStruct

	err := Struct(&st)
	if err == nil {
		t.Fatal("Expected error for invalid CAV resource name, but got none")
	}

	expectedErr := "unknown CAV resource name: invalid_resource"
	if err.Error() != expectedErr {
		t.Fatalf("Expected error message '%s', got '%s'", expectedErr, err.Error())
	}
}

func TestGenerator_CAVResourceName_Empty(t *testing.T) {
	type CStruct struct {
		Name string `fake:"{resource_name:}"`
	}

	var st CStruct

	err := Struct(&st)
	if err == nil {
		t.Fatal("Expected error for empty CAV resource name, but got none")
	}

	expectedErr := "unknown CAV resource name: "
	if err.Error() != expectedErr {
		t.Fatalf("Expected error message '%s', got '%s'", expectedErr, err.Error())
	}
}
