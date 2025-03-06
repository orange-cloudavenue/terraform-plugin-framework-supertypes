/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

// -------------------------------------------------------------------- //
// !      DO NOT EDIT. This file is auto-generated from template      ! //
// -------------------------------------------------------------------- //
package supertypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.SetValuable = SetValue{}
var _ basetypes.SetValuable = SetValueOf[struct{}]{}

// * SetType

type SetValue struct {
	basetypes.SetValue
}

func (v SetValue) Equal(o attr.Value) bool {
	other, ok := o.(SetValue)

	if !ok {
		return false
	}

	return v.SetValue.Equal(other.SetValue)
}

func (v SetValue) Type(ctx context.Context) attr.Type {
	return SetType{
		SetType: v.SetValue.Type(ctx).(basetypes.SetType),
	}
}

func (v SetValue) ToSetValue(_ context.Context) (basetypes.SetValue, diag.Diagnostics) {
	return v.SetValue, nil
}

func NewSetNull(elementType attr.Type) SetValue {
	return SetValue{
		SetValue: basetypes.NewSetNull(elementType),
	}
}

func NewSetUnknown(elementType attr.Type) SetValue {
	return SetValue{
		SetValue: basetypes.NewSetUnknown(elementType),
	}
}

func NewSetValueMust(elementType attr.Type, elements []attr.Value) SetValue {
	return SetValue{
		SetValue: basetypes.NewSetValueMust(elementType, elements),
	}
}

func NewSetValue(elementType attr.Type, elements []attr.Value) (SetValue, diag.Diagnostics) {
	x, d := basetypes.NewSetValue(elementType, elements)
	return SetValue{
		SetValue: x,
	}, d
}

func NewSetValueFrom(ctx context.Context, elementType attr.Type, elements any) (SetValue, diag.Diagnostics) {
	x, d := basetypes.NewSetValueFrom(ctx, elementType, elements)
	return SetValue{
		SetValue: x,
	}, d
}

// * CustomFunc

func (v *SetValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.SetValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *SetValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.SetValue, d = types.SetValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *SetValue) SetNull(ctx context.Context) {
	v.SetValue = basetypes.NewSetNull(v.ElementType(ctx))
}

func (v *SetValue) SetUnknown(ctx context.Context) {
	v.SetValue = basetypes.NewSetUnknown(v.ElementType(ctx))
}

func (v SetValue) IsKnown() bool {
	return !v.SetValue.IsNull() && !v.SetValue.IsUnknown()
}

// * SetTypeOf

type SetValueOf[T any] struct {
	basetypes.SetValue
}

// ToSetValue converts the given value to a SetValue.
func (v SetValueOf[T]) ToSetValue(_ context.Context) (basetypes.SetValue, diag.Diagnostics) {
	return v.SetValue, nil
}

// Equal returns true if the given value is equal to this value.
func (v SetValueOf[T]) Equal(o attr.Value) bool {
	other, ok := o.(SetValueOf[T])

	if !ok {
		return false
	}

	return v.SetValue.Equal(other.SetValue)
}

// Type returns the type of this value.
func (v SetValueOf[T]) Type(ctx context.Context) attr.Type {
	return NewSetTypeOf[T](ctx)
}

// Get returns a SetValueOf from the given value.
func (v SetValueOf[T]) Get(ctx context.Context) (values []T, diags diag.Diagnostics) {

	values = make([]T, len(v.SetValue.Elements()))

	diags.Append(v.SetValue.ElementsAs(ctx, &values, false)...)
	return
}

// MustGet returns a SetValueOf from the given value.
// It panics if set conversion fails.
func (v SetValueOf[T]) MustGet(ctx context.Context) []T {
	return MustDiag(v.Get(ctx))
}

// DiagsGet returns a SetValueOf from the given value.
// It appends diags if set conversion fails.
func (v SetValueOf[T]) DiagsGet(ctx context.Context, diags diag.Diagnostics) (values []T) {
	vv, d := v.Get(ctx)
	diags.Append(d...)
	return vv
}

// MustSet sets the value of this value.
func (v *SetValueOf[T]) MustSet(ctx context.Context, elements []T) {
	MustDiags(v.Set(ctx, elements))
}

// DiagsSet sets the value of this value.
func (v *SetValueOf[T]) DiagsSet(ctx context.Context, diags diag.Diagnostics, elements []T) {
	diags.Append(v.Set(ctx, elements)...)
}

// Set sets the value of this value.
func (v *SetValueOf[T]) Set(ctx context.Context, elements []T) diag.Diagnostics {
	var d diag.Diagnostics
	v.SetValue, d = types.SetValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

// NewSetValueOfUnknown returns a new SetValueOf with an unknown value.
func NewSetValueOfUnknown[T any](ctx context.Context) SetValueOf[T] {
	return SetValueOf[T]{
		SetValue: basetypes.NewSetUnknown(ElementTypeMust[T](ctx)),
	}
}

// NewSetValueOfNull returns a new SetValueOf with a null value.
func NewSetValueOfNull[T any](ctx context.Context) SetValueOf[T] {
	return SetValueOf[T]{
		SetValue: basetypes.NewSetNull(ElementTypeMust[T](ctx)),
	}
}

// newSetValueOf is a helper function to create a new SetValueOf.
func newSetValueOf[T any](ctx context.Context, elements any) SetValueOf[T] {
	return SetValueOf[T]{SetValue: MustDiag(basetypes.NewSetValueFrom(ctx, ElementTypeMust[T](ctx), elements))}
}

// NewSetValueOfSlice returns a new SetValueOf with the given slice value.
func NewSetValueOfSlice[T any](ctx context.Context, elements []T) SetValueOf[T] {
	return newSetValueOf[T](ctx, elements)
}

// NewSetValueOfSlicePtr returns a new SetValueOf with the given slice value.
func NewSetValueOfSlicePtr[T any](ctx context.Context, elements []*T) SetValueOf[T] {
	return newSetValueOf[T](ctx, elements)
}

// IsKnown returns true if the value is known.
func (v SetValueOf[T]) IsKnown() bool {
	return !v.SetValue.IsNull() && !v.SetValue.IsUnknown()
}

// SetNull sets the value to null.
func (v *SetValueOf[T]) SetNull(ctx context.Context) {
	(*v) = NewSetValueOfNull[T](ctx)
}

// SetUnknown sets the value to unknown.
func (v *SetValueOf[T]) SetUnknown(ctx context.Context) {
	(*v) = NewSetValueOfUnknown[T](ctx)
}
