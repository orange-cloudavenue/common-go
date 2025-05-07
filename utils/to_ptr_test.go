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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleToPTR() {
	// Example usage of ToPTR function
	str := "Hello, World!"
	ptr := ToPTR(str)
	println(reflect.TypeOf(ptr).Kind()) // Output: reflect.Ptr
}

func TestToPTR(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		val := 42
		ptr := utils.ToPTR(val)
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})

	t.Run("string", func(t *testing.T) {
		val := "hello"
		ptr := utils.ToPTR(val)
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})

	t.Run("struct", func(t *testing.T) {
		type Sample struct {
			Field string
		}
		val := Sample{Field: "test"}
		ptr := utils.ToPTR(val)
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})

	t.Run("nilable type", func(t *testing.T) {
		var val *int
		ptr := utils.ToPTR(val)
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func ExampleToPTRSlice() {
	// Example usage of ToPTRSlice function
	strs := []string{"Hello", "World"}
	ptrs := ToPTRSlice(strs)
	for _, ptr := range ptrs {
		println(reflect.TypeOf(ptr).Kind()) // Output: reflect.Ptr
	}
}

func TestToPTRSlice(t *testing.T) {
	t.Parallel()

	t.Run("int slice", func(t *testing.T) {
		vals := []int{1, 2, 3}
		ptrs := ToPTRSlice(vals)
		assert.NotNil(t, ptrs)
		assert.Equal(t, len(vals), len(ptrs))
		for i, ptr := range ptrs {
			assert.NotNil(t, ptr)
			assert.Equal(t, vals[i], *ptr)
		}
	})

	t.Run("string slice", func(t *testing.T) {
		vals := []string{"a", "b", "c"}
		ptrs := ToPTRSlice(vals)
		assert.NotNil(t, ptrs)
		assert.Equal(t, len(vals), len(ptrs))
		for i, ptr := range ptrs {
			assert.NotNil(t, ptr)
			assert.Equal(t, vals[i], *ptr)
		}
	})

	t.Run("struct slice", func(t *testing.T) {
		type Sample struct {
			Field string
		}
		vals := []Sample{{Field: "x"}, {Field: "y"}}
		ptrs := ToPTRSlice(vals)
		assert.NotNil(t, ptrs)
		assert.Equal(t, len(vals), len(ptrs))
		for i, ptr := range ptrs {
			assert.NotNil(t, ptr)
			assert.Equal(t, vals[i], *ptr)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		vals := []int{}
		ptrs := ToPTRSlice(vals)
		assert.Nil(t, ptrs)
	})

	t.Run("nil slice", func(t *testing.T) {
		var vals []int
		ptrs := ToPTRSlice(vals)
		assert.Nil(t, ptrs)
	})
}
