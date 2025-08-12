/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package validators_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/orange-cloudavenue/common-go/validators"
)

func TestCustomValidators(t *testing.T) {
	t.Parallel()
	testsCase := map[string]struct {
		valuesWork        []any
		valuesDoesNotWork []any
		rule              string
	}{
		"cases-disallow_upper-upper": {
			valuesWork:        []any{"test"},
			valuesDoesNotWork: []any{"Test"},
			rule:              "disallow_upper",
		},
		"cases-disallow_space-space": {
			valuesWork:        []any{"testtest"},
			valuesDoesNotWork: []any{"test test!"},
			rule:              "disallow_space",
		},
		"http_status_code": {
			valuesWork:        []any{"200", 200, 404},
			valuesDoesNotWork: []any{"666", "two hundred", 666},
			rule:              "http_status_code",
		},
		"http_status_code_range": {
			valuesWork:        []any{"200-404", "200-299"},
			valuesDoesNotWork: []any{"10-200", "100-20", "200-100", "200-404-500", "200-invalid", "invalid-404"},
			rule:              "http_status_code_range",
		},
		"urn": {
			valuesWork:        []any{"urn:vcloud:gateway:4aeb40d8-038c-4e77-8181-a7054f583b12"},
			valuesDoesNotWork: []any{"urn:vcloud:vm:invalid"},
			rule:              "urn=edgegateway",
		},
		"urn-bad": {
			valuesWork:        []any{},
			valuesDoesNotWork: []any{"urn:vcloud:gateway:4aeb40d8-038c-4e77-8181-a7054f583b12"},
			rule:              "urn=bad",
		},
		"ipv4_range": {
			valuesWork:        []any{"192.168.0.1-192.168.0.100"},
			valuesDoesNotWork: []any{"192.168.0.256-192.168.0.300", "192.168.0.256", "192.168.0.100-192.168.0.1"},
			rule:              "ipv4_range",
		},
		"tcp_udp_port": {
			valuesWork:        []any{"80", 80, "65535", 65535},
			valuesDoesNotWork: []any{"-1", "65536", "invalid", 65536, "", 0},
			rule:              "tcp_udp_port",
		},
		"tcp_udp_port_range": {
			valuesWork:        []any{"80-100", "80-65535"},
			valuesDoesNotWork: []any{"-1", "65536", "invalid", "", "100-80", "100-65538", "65538-100", "invalid-80", "80-invalid"},
			rule:              "tcp_udp_port_range",
		},
		"str_key_value": {
			valuesWork:        []any{"key=value", "key=val"},
			valuesDoesNotWork: []any{"key=value=val", "key=val=val", "key=val=val=val"},
			rule:              "str_key_value",
		},
		"cav_resource_name": {
			valuesWork:        []any{"tn01e02ocb0001234spt101", "tn01e02ocb0001234spt102"},
			valuesDoesNotWork: []any{"prvrf01eocb0001234allsp01", "invalid"},
			rule:              "resource_name=edgegateway",
		},
		"case-camelCase": {
			valuesWork:        []any{"camelCaseExample", "anotherCamelCase"},
			valuesDoesNotWork: []any{"CamelCase", "camel_case", "kebab-case", "UPPER_CASE"},
			rule:              "case=camelCase",
		},
		"case-snake_case": {
			valuesWork:        []any{"snake_case_example", "another_snake_case"},
			valuesDoesNotWork: []any{"snakeCase", "kebab-case", "UPPER_CASE", "CamelCase"},
			rule:              "case=snake_case",
		},
		"case-PascalCase": {
			valuesWork:        []any{"PascalCaseExample", "AnotherPascalCase"},
			valuesDoesNotWork: []any{"pascalCase", "snake_case", "kebab-case", "UPPER_CASE"},
			rule:              "case=PascalCase",
		},
		"case-UPPER_CASE": {
			valuesWork:        []any{"UPPER_CASE_EXAMPLE", "ANOTHER_UPPER_CASE"},
			valuesDoesNotWork: []any{"upper_case", "kebab-case", "camelCase", "PascalCase"},
			rule:              "case=UPPER_CASE",
		},
		"case-kebab-case": {
			valuesWork:        []any{"kebab-case-example", "another-kebab-case"},
			valuesDoesNotWork: []any{"kebabCase", "snake_case", "UPPER_CASE", "PascalCase"},
			rule:              "case=kebab-case",
		},
		"case-invalid-format": {
			valuesWork:        []any{},
			valuesDoesNotWork: []any{"invalidCaseFormat"},
			rule:              "case=invalidFormat",
		},
	}

	for name, test := range testsCase {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			for _, value := range test.valuesWork {
				if err := validators.New().Var(value, test.rule); err != nil {
					t.Errorf("expected no error for value: %v, rule: %s", value, test.rule)
				}
			}

			for _, value := range test.valuesDoesNotWork {
				if err := validators.New().Var(value, test.rule); err == nil {
					t.Errorf("expected error for value: %v, rule: %s", value, test.rule)
				}
			}
		})
	}
}

func TestDefaulter(t *testing.T) {
	t.Parallel()
	type defaultTest struct {
		Field1      string  `default:"default_value"`
		Field2      int     `default:"-42"`
		Field3      bool    `default:"true"`
		Field4      float64 `default:"3.14"`
		Field5      uint64  `default:"1000"`
		FieldStruct struct {
			SubField1 string `default:"sub_default_value"`
			SubField2 int    `default:"100"`
		}
	}

	v := validators.New()
	defaults := defaultTest{}
	if err := v.Struct(&defaults); err != nil {
		t.Errorf("expected no error when setting defaults: %v", err)
	}

	if err := v.StructCtx(t.Context(), &defaults); err != nil {
		t.Errorf("expected no error when setting defaults with context: %v", err)
	}

	assert.Equal(t, "default_value", defaults.Field1, "Field1 should have default value")
	assert.Equal(t, -42, defaults.Field2, "Field2 should have default value")
	assert.Equal(t, true, defaults.Field3, "Field3 should have default value")
	assert.Equal(t, 3.14, defaults.Field4, "Field4 should have default value")
	assert.Equal(t, uint64(1000), defaults.Field5, "Field5 should have default value")
	assert.Equal(t, "sub_default_value", defaults.FieldStruct.SubField1, "SubField1 should have default value")
	assert.Equal(t, 100, defaults.FieldStruct.SubField2, "SubField2 should have default value")
}
