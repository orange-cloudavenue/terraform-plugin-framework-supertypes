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
var _ basetypes.ListValuable = ListValue{}
var _ basetypes.ListValuable = ListValueOf[struct{}]{}

// * ListType

type ListValue struct {
	basetypes.ListValue
}

func (v ListValue) Equal(o attr.Value) bool {
	other, ok := o.(ListValue)

	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

func (v ListValue) Type(ctx context.Context) attr.Type {
	return ListType{
		ListType: v.ListValue.Type(ctx).(basetypes.ListType),
	}
}

func (v ListValue) ToListValue(_ context.Context) (basetypes.ListValue, diag.Diagnostics) {
	return v.ListValue, nil
}

func NewListNull(elementType attr.Type) ListValue {
	return ListValue{
		ListValue: basetypes.NewListNull(elementType),
	}
}

func NewListUnknown(elementType attr.Type) ListValue {
	return ListValue{
		ListValue: basetypes.NewListUnknown(elementType),
	}
}

func NewListValueMust(elementType attr.Type, elements []attr.Value) ListValue {
	return ListValue{
		ListValue: basetypes.NewListValueMust(elementType, elements),
	}
}

func NewListValue(elementType attr.Type, elements []attr.Value) (ListValue, diag.Diagnostics) {
	x, d := basetypes.NewListValue(elementType, elements)
	return ListValue{
		ListValue: x,
	}, d
}

func NewListValueFrom(ctx context.Context, elementType attr.Type, elements any) (ListValue, diag.Diagnostics) {
	x, d := basetypes.NewListValueFrom(ctx, elementType, elements)
	return ListValue{
		ListValue: x,
	}, d
}

// * CustomFunc

func (v *ListValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.ListValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *ListValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.ListValue, d = types.ListValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *ListValue) SetNull(ctx context.Context) {
	v.ListValue = basetypes.NewListNull(v.ElementType(ctx))
}

func (v *ListValue) SetUnknown(ctx context.Context) {
	v.ListValue = basetypes.NewListUnknown(v.ElementType(ctx))
}

func (v ListValue) IsKnown() bool {
	return !v.ListValue.IsNull() && !v.ListValue.IsUnknown()
}

// * ListTypeOf

type ListValueOf[T any] struct {
	basetypes.ListValue
}

// ToListValue converts the given value to a ListValue.
func (v ListValueOf[T]) ToListValue(_ context.Context) (basetypes.ListValue, diag.Diagnostics) {
	return v.ListValue, nil
}

// Equal returns true if the given value is equal to this value.
func (v ListValueOf[T]) Equal(o attr.Value) bool {
	other, ok := o.(ListValueOf[T])

	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

// Type returns the type of this value.
func (v ListValueOf[T]) Type(ctx context.Context) attr.Type {
	return NewListTypeOf[T](ctx)
}

// Get returns a ListValueOf from the given value.
func (v ListValueOf[T]) Get(ctx context.Context) (values []T, diags diag.Diagnostics) {

	values = make([]T, len(v.ListValue.Elements()))

	diags.Append(v.ListValue.ElementsAs(ctx, &values, false)...)
	return
}

// MustGet returns a ListValueOf from the given value.
// It panics if set conversion fails.
func (v ListValueOf[T]) MustGet(ctx context.Context) []T {
	return MustDiag(v.Get(ctx))
}

// DiagsGet returns a ListValueOf from the given value.
// It appends diags if set conversion fails.
func (v ListValueOf[T]) DiagsGet(ctx context.Context, diags diag.Diagnostics) (values []T) {
	vv, d := v.Get(ctx)
	diags.Append(d...)
	return vv
}

// MustSet sets the value of this value.
func (v *ListValueOf[T]) MustSet(ctx context.Context, elements []T) {
	MustDiags(v.Set(ctx, elements))
}

// DiagsSet sets the value of this value.
func (v *ListValueOf[T]) DiagsSet(ctx context.Context, diags diag.Diagnostics, elements []T) {
	diags.Append(v.Set(ctx, elements)...)
}

// Set sets the value of this value.
func (v *ListValueOf[T]) Set(ctx context.Context, elements []T) diag.Diagnostics {
	var d diag.Diagnostics
	v.ListValue, d = types.ListValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

// NewListValueOfUnknown returns a new ListValueOf with an unknown value.
func NewListValueOfUnknown[T any](ctx context.Context) ListValueOf[T] {
	return ListValueOf[T]{
		ListValue: basetypes.NewListUnknown(ElementTypeMust[T](ctx)),
	}
}

// NewListValueOfNull returns a new ListValueOf with a null value.
func NewListValueOfNull[T any](ctx context.Context) ListValueOf[T] {
	return ListValueOf[T]{
		ListValue: basetypes.NewListNull(ElementTypeMust[T](ctx)),
	}
}

// newListValueOf is a helper function to create a new ListValueOf.
func newListValueOf[T any](ctx context.Context, elements any) ListValueOf[T] {
	return ListValueOf[T]{ListValue: MustDiag(basetypes.NewListValueFrom(ctx, ElementTypeMust[T](ctx), elements))}
}

// NewListValueOfSlice returns a new ListValueOf with the given slice value.
func NewListValueOfSlice[T any](ctx context.Context, elements []T) ListValueOf[T] {
	return newListValueOf[T](ctx, elements)
}

// NewListValueOfSlicePtr returns a new ListValueOf with the given slice value.
func NewListValueOfSlicePtr[T any](ctx context.Context, elements []*T) ListValueOf[T] {
	return newListValueOf[T](ctx, elements)
}

// IsKnown returns true if the value is known.
func (v ListValueOf[T]) IsKnown() bool {
	return !v.ListValue.IsNull() && !v.ListValue.IsUnknown()
}

// SetNull sets the value to null.
func (v *ListValueOf[T]) SetNull(ctx context.Context) {
	(*v) = NewListValueOfNull[T](ctx)
}

// SetUnknown sets the value to unknown.
func (v *ListValueOf[T]) SetUnknown(ctx context.Context) {
	(*v) = NewListValueOfUnknown[T](ctx)
}
