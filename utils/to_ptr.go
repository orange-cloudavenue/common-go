/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package utils

// ToPTR is a generic function that takes a value and returns a pointer to it.
func ToPTR[T any](v T) *T {
	return &v
}

// ToPTRSlice is a generic function that takes a slice of values and returns a slice of pointers to them.
func ToPTRSlice[T any](v []T) []*T {
	if len(v) == 0 {
		return nil
	}

	slicePointer := make([]*T, len(v))
	for i, item := range v {
		// FIX variable into the local variable for bypassing the pointer aliasing (https://stackoverflow.com/a/64715804)
		item := item //nolint:copyloopvar
		slicePointer[i] = &item
	}
	return slicePointer
}
