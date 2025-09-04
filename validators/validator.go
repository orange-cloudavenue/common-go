/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package validators

import (
	"context"
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	*validator.Validate
}

// New creates a new validator.
func New() *Validator {
	// * String
	v := validator.New(validator.WithRequiredStructEnabled())
	_ = v.RegisterValidation(DisallowUpper.Key, DisallowUpper.Func)
	_ = v.RegisterValidation(DisallowSpace.Key, DisallowSpace.Func)
	_ = v.RegisterValidation(Case.Key, Case.Func)

	// * Key/Value
	_ = v.RegisterValidation(KeyValue.Key, KeyValue.Func)

	// * Cloud Avenue
	_ = v.RegisterValidation(URN.Key, URN.Func)
	_ = v.RegisterValidation(CAVResourceName.Key, CAVResourceName.Func)

	// * Network
	_ = v.RegisterValidation(IPV4Range.Key, IPV4Range.Func)
	_ = v.RegisterValidation(TCPUDPPort.Key, TCPUDPPort.Func)
	_ = v.RegisterValidation(TCPUDPPortRange.Key, TCPUDPPortRange.Func)

	// * HTTP
	_ = v.RegisterValidation(HTTPStatusCode.Key, HTTPStatusCode.Func)
	_ = v.RegisterValidation(HTTPStatusCodeRange.Key, HTTPStatusCodeRange.Func)

	// * Require/Exclude
	_ = v.RegisterValidation(RequireIfNull.Key, RequireIfNull.Func)
	_ = v.RegisterValidation(ExcludeIfNull.Key, ExcludeIfNull.Func)

	return &Validator{
		Validate: v,
	}
}

func (v *Validator) Struct(s interface{}) error {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		return errors.New("validator: Struct() expects a pointer to a struct")
	}

	if err := v.defaulter(s); err != nil {
		return err
	}

	if err := v.Validate.Struct(s); err != nil {
		return err
	}

	return nil
}

func (v *Validator) StructCtx(ctx context.Context, s interface{}) error {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		return errors.New("validator: StructCtx() expects a pointer to a struct")
	}

	if err := v.defaulter(s); err != nil {
		return err
	}

	if err := v.Validate.StructCtx(ctx, s); err != nil {
		return err
	}

	return nil
}
