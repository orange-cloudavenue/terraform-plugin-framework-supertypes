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
	_ basetypes.SetValuable = SetValue{}

	_ basetypes.SetValuable = SetNestedObjectValueOf[struct{}]{}
	_ NestedObjectValue     = SetNestedObjectValueOf[struct{}]{}
)

type SetNestedValue struct {
	basetypes.SetValue
}

func NewSetNestedObjectTypeOf[T any](ctx context.Context) SetNestedObjectTypeOf[T] {
	return SetNestedObjectTypeOf[T]{basetypes.SetType{ElemType: NewObjectTypeOf[T](ctx)}}
}

// SetNestedObjectValueOf represents a Terraform Plugin Framework Set value whose elements are of type ObjectTypeOf.
type SetNestedObjectValueOf[T any] struct {
	basetypes.SetValue
}

func (v SetNestedValue) Type(ctx context.Context) attr.Type {
	return v.SetValue.Type(ctx)
}

func (v SetNestedObjectValueOf[T]) Type(ctx context.Context) attr.Type {
	return NewSetNestedObjectTypeOf[T](ctx)
}

func (v SetNestedObjectValueOf[T]) Equal(o attr.Value) bool {
	other, ok := o.(SetNestedObjectValueOf[T])

	if !ok {
		return false
	}

	return v.SetValue.Equal(other.SetValue)
}

func (v SetNestedValue) ToSetValue(_ context.Context) (basetypes.SetValue, diag.Diagnostics) {
	return v.SetValue, nil
}

func NewSetNestedNull(elementType attr.Type) SetNestedValue {
	return SetNestedValue{
		SetValue: basetypes.NewSetNull(elementType),
	}
}

func NewSetNestedUnknown(elementType attr.Type) SetNestedValue {
	return SetNestedValue{
		SetValue: basetypes.NewSetUnknown(elementType),
	}
}

// * CustomFunc

func (v *SetNestedValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.SetValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *SetNestedValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.SetValue, d = types.SetValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *SetNestedValue) SetNull(ctx context.Context) {
	v.SetValue = basetypes.NewSetNull(v.ElementType(ctx))
}

func (v *SetNestedValue) SetUnknown(ctx context.Context) {
	v.SetValue = basetypes.NewSetUnknown(v.ElementType(ctx))
}

func (v SetNestedValue) IsKnown() bool {
	return !v.SetValue.IsNull() && !v.SetValue.IsUnknown()
}

// Get returns a slice of pointers to the elements of a SetNestedObject.
func (v SetNestedObjectValueOf[T]) Get(ctx context.Context) ([]*T, diag.Diagnostics) {
	return nestedObjectValueSlice[T](ctx, v.SetValue)
}

// Set returns a SetNestedObjectValueOf from a slice of pointers to the elements of a SetNestedObject.
func (v *SetNestedObjectValueOf[T]) Set(ctx context.Context, slice []*T) diag.Diagnostics {
	var diags diag.Diagnostics
	v.SetValue, diags = basetypes.NewSetValueFrom(ctx, NewObjectTypeOf[T](ctx), slice)
	return diags
}
func NewSetNestedObjectValueOfNull[T any](ctx context.Context) SetNestedObjectValueOf[T] {
	return SetNestedObjectValueOf[T]{SetValue: basetypes.NewSetNull(NewObjectTypeOf[T](ctx))}
}

func NewSetNestedObjectValueOfUnknown[T any](ctx context.Context) SetNestedObjectValueOf[T] {
	return SetNestedObjectValueOf[T]{SetValue: basetypes.NewSetUnknown(NewObjectTypeOf[T](ctx))}
}

func NewSetNestedObjectValueOfPtr[T any](ctx context.Context, t *T) SetNestedObjectValueOf[T] {
	return NewSetNestedObjectValueOfSlice(ctx, []*T{t})
}
func NewSetNestedObjectValueOfSlice[T any](ctx context.Context, ts []*T) SetNestedObjectValueOf[T] {
	return newSetNestedObjectValueOf[T](ctx, ts)
}
func NewSetNestedObjectValueOfValueSlice[T any](ctx context.Context, ts []T) SetNestedObjectValueOf[T] {
	return newSetNestedObjectValueOf[T](ctx, ts)
}
func newSetNestedObjectValueOf[T any](ctx context.Context, elements any) SetNestedObjectValueOf[T] {
	return SetNestedObjectValueOf[T]{SetValue: MustDiag(basetypes.NewSetValueFrom(ctx, NewObjectTypeOf[T](ctx), elements))}
}

// IsKnown returns whether the value is known.
func (v SetNestedObjectValueOf[T]) IsKnown() bool {
	if !v.IsNull() && !v.IsUnknown() {
		return true
	}

	return false
}

func (v *SetNestedObjectValueOf[T]) SetNull(ctx context.Context) {
	*v = NewSetNestedObjectValueOfNull[T](ctx)
}

func (v *SetNestedObjectValueOf[T]) SetUnknown(ctx context.Context) {
	*v = NewSetNestedObjectValueOfUnknown[T](ctx)
}
