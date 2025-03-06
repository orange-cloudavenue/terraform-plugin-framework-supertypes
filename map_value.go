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
var _ basetypes.MapValuable = MapValue{}
var _ basetypes.MapValuable = MapValueOf[struct{}]{}

// * MapType

type MapValue struct {
	basetypes.MapValue
}

func (v MapValue) Equal(o attr.Value) bool {
	other, ok := o.(MapValue)

	if !ok {
		return false
	}

	return v.MapValue.Equal(other.MapValue)
}

func (v MapValue) Type(ctx context.Context) attr.Type {
	return MapType{
		MapType: v.MapValue.Type(ctx).(basetypes.MapType),
	}
}

func (v MapValue) ToMapValue(_ context.Context) (basetypes.MapValue, diag.Diagnostics) {
	return v.MapValue, nil
}

func NewMapNull(elementType attr.Type) MapValue {
	return MapValue{
		MapValue: basetypes.NewMapNull(elementType),
	}
}

func NewMapUnknown(elementType attr.Type) MapValue {
	return MapValue{
		MapValue: basetypes.NewMapUnknown(elementType),
	}
}

func NewMapValueMust(elementType attr.Type, elements map[string]attr.Value) MapValue {
	return MapValue{
		MapValue: basetypes.NewMapValueMust(elementType, elements),
	}
}

func NewMapValue(elementType attr.Type, elements map[string]attr.Value) (MapValue, diag.Diagnostics) {
	x, d := basetypes.NewMapValue(elementType, elements)
	return MapValue{
		MapValue: x,
	}, d
}

func NewMapValueFrom(ctx context.Context, elementType attr.Type, elements any) (MapValue, diag.Diagnostics) {
	x, d := basetypes.NewMapValueFrom(ctx, elementType, elements)
	return MapValue{
		MapValue: x,
	}, d
}

// * CustomFunc

func (v *MapValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.MapValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *MapValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.MapValue, d = types.MapValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *MapValue) SetNull(ctx context.Context) {
	v.MapValue = basetypes.NewMapNull(v.ElementType(ctx))
}

func (v *MapValue) SetUnknown(ctx context.Context) {
	v.MapValue = basetypes.NewMapUnknown(v.ElementType(ctx))
}

func (v MapValue) IsKnown() bool {
	return !v.MapValue.IsNull() && !v.MapValue.IsUnknown()
}

// * MapTypeOf

type MapValueOf[T any] struct {
	basetypes.MapValue
}

// ToMapValue converts the given value to a MapValue.
func (v MapValueOf[T]) ToMapValue(_ context.Context) (basetypes.MapValue, diag.Diagnostics) {
	return v.MapValue, nil
}

// Equal returns true if the given value is equal to this value.
func (v MapValueOf[T]) Equal(o attr.Value) bool {
	other, ok := o.(MapValueOf[T])

	if !ok {
		return false
	}

	return v.MapValue.Equal(other.MapValue)
}

// Type returns the type of this value.
func (v MapValueOf[T]) Type(ctx context.Context) attr.Type {
	return NewMapTypeOf[T](ctx)
}

// Get returns a MapValueOf from the given value.
func (v MapValueOf[T]) Get(ctx context.Context) (values map[string]T, diags diag.Diagnostics) {

	values = make(map[string]T, len(v.MapValue.Elements()))

	diags.Append(v.MapValue.ElementsAs(ctx, &values, false)...)
	return
}

// MustGet returns a MapValueOf from the given value.
// It panics if set conversion fails.
func (v MapValueOf[T]) MustGet(ctx context.Context) map[string]T {
	return MustDiag(v.Get(ctx))
}

// DiagsGet returns a MapValueOf from the given value.
// It appends diags if set conversion fails.
func (v MapValueOf[T]) DiagsGet(ctx context.Context, diags diag.Diagnostics) (values map[string]T) {
	vv, d := v.Get(ctx)
	diags.Append(d...)
	return vv
}

// MustSet sets the value of this value.
func (v *MapValueOf[T]) MustSet(ctx context.Context, elements map[string]T) {
	MustDiags(v.Set(ctx, elements))
}

// DiagsSet sets the value of this value.
func (v *MapValueOf[T]) DiagsSet(ctx context.Context, diags diag.Diagnostics, elements map[string]T) {
	diags.Append(v.Set(ctx, elements)...)
}

// Set sets the value of this value.
func (v *MapValueOf[T]) Set(ctx context.Context, elements map[string]T) diag.Diagnostics {
	var d diag.Diagnostics
	v.MapValue, d = types.MapValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

// NewMapValueOfUnknown returns a new MapValueOf with an unknown value.
func NewMapValueOfUnknown[T any](ctx context.Context) MapValueOf[T] {
	return MapValueOf[T]{
		MapValue: basetypes.NewMapUnknown(ElementTypeMust[T](ctx)),
	}
}

// NewMapValueOfNull returns a new MapValueOf with a null value.
func NewMapValueOfNull[T any](ctx context.Context) MapValueOf[T] {
	return MapValueOf[T]{
		MapValue: basetypes.NewMapNull(ElementTypeMust[T](ctx)),
	}
}

// newMapValueOf is a helper function to create a new MapValueOf.
func newMapValueOf[T any](ctx context.Context, elements any) MapValueOf[T] {
	return MapValueOf[T]{
		MapValue: MustDiag(types.MapValueFrom(ctx, ElementTypeMust[T](ctx), elements)),
	}
}

// NewMapValueOfMap returns a new MapValueOf with the given map value.
func NewMapValueOfMap[T any](ctx context.Context, elements map[string]T) (MapValueOf[T], diag.Diagnostics) {
	v, d := types.MapValueFrom(ctx, ElementTypeMust[T](ctx), elements)
	return MapValueOf[T]{
		MapValue: v,
	}, d
}

// IsKnown returns true if the value is known.
func (v MapValueOf[T]) IsKnown() bool {
	return !v.MapValue.IsNull() && !v.MapValue.IsUnknown()
}

// SetNull sets the value to null.
func (v *MapValueOf[T]) SetNull(ctx context.Context) {
	(*v) = NewMapValueOfNull[T](ctx)
}

// SetUnknown sets the value to unknown.
func (v *MapValueOf[T]) SetUnknown(ctx context.Context) {
	(*v) = NewMapValueOfUnknown[T](ctx)
}
