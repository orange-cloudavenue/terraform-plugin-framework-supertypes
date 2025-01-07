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
var (
	_ basetypes.MapValuable = MapValue{}

	_ basetypes.MapValuable = MapNestedObjectValueOf[struct{}]{}
	_ NestedObjectValue     = MapNestedObjectValueOf[struct{}]{}
)

type MapNestedValue struct {
	basetypes.MapValue
}

func NewMapNestedObjectTypeOf[T any](ctx context.Context) MapNestedObjectTypeOf[T] {
	return MapNestedObjectTypeOf[T]{basetypes.MapType{ElemType: NewObjectTypeOf[T](ctx)}}
}

// MapNestedObjectValueOf represents a Terraform Plugin Framework Map value whose elements are of type ObjectTypeOf.
type MapNestedObjectValueOf[T any] struct {
	basetypes.MapValue
}

func (v MapNestedValue) Type(ctx context.Context) attr.Type {
	return v.MapValue.Type(ctx)
}

func (v MapNestedObjectValueOf[T]) Type(ctx context.Context) attr.Type {
	return NewMapNestedObjectTypeOf[T](ctx)
}

func (v MapNestedObjectValueOf[T]) Equal(o attr.Value) bool {
	other, ok := o.(MapNestedObjectValueOf[T])

	if !ok {
		return false
	}

	return v.MapValue.Equal(other.MapValue)
}

func (v MapNestedValue) ToMapValue(_ context.Context) (basetypes.MapValue, diag.Diagnostics) {
	return v.MapValue, nil
}

func NewMapNestedNull(elementType attr.Type) MapNestedValue {
	return MapNestedValue{
		MapValue: basetypes.NewMapNull(elementType),
	}
}

func NewMapNestedUnknown(elementType attr.Type) MapNestedValue {
	return MapNestedValue{
		MapValue: basetypes.NewMapUnknown(elementType),
	}
}

// * CustomFunc

func (v *MapNestedValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.MapValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *MapNestedValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.MapValue, d = types.MapValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *MapNestedValue) SetNull(ctx context.Context) {
	v.MapValue = basetypes.NewMapNull(v.ElementType(ctx))
}

func (v *MapNestedValue) SetUnknown(ctx context.Context) {
	v.MapValue = basetypes.NewMapUnknown(v.ElementType(ctx))
}

func (v MapNestedValue) IsKnown() bool {
	return !v.MapValue.IsNull() && !v.MapValue.IsUnknown()
}

// Get returns a slice of pointers to the elements of a MapNestedObject.
func (v MapNestedObjectValueOf[T]) Get(ctx context.Context) (map[string]*T, diag.Diagnostics) {
	return nestedObjectValueMap[T](ctx, v.MapValue)
}

// Set returns a MapNestedObjectValueOf from a slice of pointers to the elements of a MapNestedObject.
func (v *MapNestedObjectValueOf[T]) Set(ctx context.Context, m map[string]*T) diag.Diagnostics {
	var diags diag.Diagnostics
	v.MapValue, diags = basetypes.NewMapValueFrom(ctx, NewObjectTypeOf[T](ctx), m)
	return diags
}
func NewMapNestedObjectValueOfNull[T any](ctx context.Context) MapNestedObjectValueOf[T] {
	return MapNestedObjectValueOf[T]{MapValue: basetypes.NewMapNull(NewObjectTypeOf[T](ctx))}
}

func NewMapNestedObjectValueOfUnknown[T any](ctx context.Context) MapNestedObjectValueOf[T] {
	return MapNestedObjectValueOf[T]{MapValue: basetypes.NewMapUnknown(NewObjectTypeOf[T](ctx))}
}

func NewMapNestedObjectValueOfPtr[T any](ctx context.Context, m map[string]*T) MapNestedObjectValueOf[T] {
	return NewMapNestedObjectValueOfMap(ctx, m)
}
func NewMapNestedObjectValueOfMap[T any](ctx context.Context, m map[string]*T) MapNestedObjectValueOf[T] {
	return newMapNestedObjectValueOf[T](ctx, m)
}
func NewMapNestedObjectValueOfValueMap[T any](ctx context.Context, m map[string]T) MapNestedObjectValueOf[T] {
	return newMapNestedObjectValueOf[T](ctx, m)
}
func newMapNestedObjectValueOf[T any](ctx context.Context, elements any) MapNestedObjectValueOf[T] {
	return MapNestedObjectValueOf[T]{MapValue: MustDiag(basetypes.NewMapValueFrom(ctx, NewObjectTypeOf[T](ctx), elements))}
}

// IsKnown returns whether the value is known.
func (v MapNestedObjectValueOf[T]) IsKnown() bool {
	if !v.IsNull() && !v.IsUnknown() {
		return true
	}

	return false
}

func (v *MapNestedObjectValueOf[T]) SetNull(ctx context.Context) {
	*v = NewMapNestedObjectValueOfNull[T](ctx)
}

func (v *MapNestedObjectValueOf[T]) SetUnknown(ctx context.Context) {
	*v = NewMapNestedObjectValueOfUnknown[T](ctx)
}
