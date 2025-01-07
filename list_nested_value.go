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
	_ basetypes.ListValuable = ListValue{}

	_ basetypes.ListValuable = ListNestedObjectValueOf[struct{}]{}
	_ NestedObjectValue      = ListNestedObjectValueOf[struct{}]{}
)

type ListNestedValue struct {
	basetypes.ListValue
}

func NewListNestedObjectTypeOf[T any](ctx context.Context) ListNestedObjectTypeOf[T] {
	return ListNestedObjectTypeOf[T]{basetypes.ListType{ElemType: NewObjectTypeOf[T](ctx)}}
}

// ListNestedObjectValueOf represents a Terraform Plugin Framework List value whose elements are of type ObjectTypeOf.
type ListNestedObjectValueOf[T any] struct {
	basetypes.ListValue
}

func (v ListNestedValue) Type(ctx context.Context) attr.Type {
	return v.ListValue.Type(ctx)
}

func (v ListNestedObjectValueOf[T]) Type(ctx context.Context) attr.Type {
	return NewListNestedObjectTypeOf[T](ctx)
}

func (v ListNestedObjectValueOf[T]) Equal(o attr.Value) bool {
	other, ok := o.(ListNestedObjectValueOf[T])

	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

func (v ListNestedValue) ToListValue(_ context.Context) (basetypes.ListValue, diag.Diagnostics) {
	return v.ListValue, nil
}

func NewListNestedNull(elementType attr.Type) ListNestedValue {
	return ListNestedValue{
		ListValue: basetypes.NewListNull(elementType),
	}
}

func NewListNestedUnknown(elementType attr.Type) ListNestedValue {
	return ListNestedValue{
		ListValue: basetypes.NewListUnknown(elementType),
	}
}

// * CustomFunc

func (v *ListNestedValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.ListValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *ListNestedValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.ListValue, d = types.ListValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *ListNestedValue) SetNull(ctx context.Context) {
	v.ListValue = basetypes.NewListNull(v.ElementType(ctx))
}

func (v *ListNestedValue) SetUnknown(ctx context.Context) {
	v.ListValue = basetypes.NewListUnknown(v.ElementType(ctx))
}

func (v ListNestedValue) IsKnown() bool {
	return !v.ListValue.IsNull() && !v.ListValue.IsUnknown()
}

// Get returns a slice of pointers to the elements of a ListNestedObject.
func (v ListNestedObjectValueOf[T]) Get(ctx context.Context) ([]*T, diag.Diagnostics) {
	return nestedObjectValueSlice[T](ctx, v.ListValue)
}

// Set returns a ListNestedObjectValueOf from a slice of pointers to the elements of a ListNestedObject.
func (v *ListNestedObjectValueOf[T]) Set(ctx context.Context, slice []*T) diag.Diagnostics {
	var diags diag.Diagnostics
	v.ListValue, diags = basetypes.NewListValueFrom(ctx, NewObjectTypeOf[T](ctx), slice)
	return diags
}
func NewListNestedObjectValueOfNull[T any](ctx context.Context) ListNestedObjectValueOf[T] {
	return ListNestedObjectValueOf[T]{ListValue: basetypes.NewListNull(NewObjectTypeOf[T](ctx))}
}

func NewListNestedObjectValueOfUnknown[T any](ctx context.Context) ListNestedObjectValueOf[T] {
	return ListNestedObjectValueOf[T]{ListValue: basetypes.NewListUnknown(NewObjectTypeOf[T](ctx))}
}

func NewListNestedObjectValueOfPtr[T any](ctx context.Context, t *T) ListNestedObjectValueOf[T] {
	return NewListNestedObjectValueOfSlice(ctx, []*T{t})
}
func NewListNestedObjectValueOfSlice[T any](ctx context.Context, ts []*T) ListNestedObjectValueOf[T] {
	return newListNestedObjectValueOf[T](ctx, ts)
}
func NewListNestedObjectValueOfValueSlice[T any](ctx context.Context, ts []T) ListNestedObjectValueOf[T] {
	return newListNestedObjectValueOf[T](ctx, ts)
}
func newListNestedObjectValueOf[T any](ctx context.Context, elements any) ListNestedObjectValueOf[T] {
	return ListNestedObjectValueOf[T]{ListValue: MustDiag(basetypes.NewListValueFrom(ctx, NewObjectTypeOf[T](ctx), elements))}
}

// IsKnown returns whether the value is known.
func (v ListNestedObjectValueOf[T]) IsKnown() bool {
	if !v.IsNull() && !v.IsUnknown() {
		return true
	}

	return false
}

func (v *ListNestedObjectValueOf[T]) SetNull(ctx context.Context) {
	*v = NewListNestedObjectValueOfNull[T](ctx)
}

func (v *ListNestedObjectValueOf[T]) SetUnknown(ctx context.Context) {
	*v = NewListNestedObjectValueOfUnknown[T](ctx)
}
