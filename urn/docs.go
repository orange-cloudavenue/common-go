/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

// Package urn provides utilities for working with Uniform Resource Names (URNs)
// in the context of Cloud Avenue and vCloud resources.
//
// This package defines constants for all supported URN types, such as
// organizations, virtual machines, users, groups, networks, and more. It
// provides functions and methods to validate, normalize, and extract
// information from URNs, as well as to check the type of a given URN.
//
// Typical usage includes checking if a string is a valid URN of a specific
// type, extracting the UUID from a URN, or normalizing a UUID to a URN with
// the correct prefix.
//
// Example:
//
//	urn := urn.VM.String() + "12345678-1234-1234-1234-123456789012"
//	if urn.IsVM() {
//	    uuid := urn.extractUUIDv4(urn.VM)
//	    // Use the uuid...
//	}
//
// Thread safety:
// All exported functions are safe for concurrent use by multiple goroutines.
package urn
