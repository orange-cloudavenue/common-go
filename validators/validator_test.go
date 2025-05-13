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
			rule:              "urn=gateway",
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
