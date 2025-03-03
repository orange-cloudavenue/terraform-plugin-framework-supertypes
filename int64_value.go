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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.Int64Valuable = Int64Value{}

type Int64Value struct {
	basetypes.Int64Value
}

func (v Int64Value) Equal(o attr.Value) bool {
	other, ok := o.(Int64Value)

	if !ok {
		return false
	}

	return v.Int64Value.Equal(other.Int64Value)
}

func (v Int64Value) Type(ctx context.Context) attr.Type {
	return Int64Type{
		Int64Type: v.Int64Value.Type(ctx).(basetypes.Int64Type),
	}
}

func NewInt64Null() Int64Value {
	return Int64Value{
		Int64Value: basetypes.NewInt64Null(),
	}
}

func NewInt64Unknown() Int64Value {
	return Int64Value{
		Int64Value: basetypes.NewInt64Unknown(),
	}
}

func NewInt64Value(s int64) Int64Value {
	return Int64Value{
		Int64Value: basetypes.NewInt64Value(s),
	}
}

func NewInt64PointerValue(s *int64) Int64Value {
	return Int64Value{
		Int64Value: basetypes.NewInt64PointerValue(s),
	}
}

// * CustomFunc

// Get returns the known Int64 value. If Int64 is null or unknown, returns 0.
func (v *Int64Value) Get() int64 {
	return v.Int64Value.ValueInt64()
}

// GetPtr returns a pointer to the known int64 value, nil for a
// null value, or a pointer to 0 for an unknown value.
func (v *Int64Value) GetPtr() *int64 {
	return v.Int64Value.ValueInt64Pointer()
}

// Set sets the Int64 value.
func (v *Int64Value) Set(s int64) {

	v.Int64Value = basetypes.NewInt64Value(s)
}

// SetPtr sets a pointer to the Int64 value.
func (v *Int64Value) SetPtr(s *int64) {
	if s == nil {
		v.Int64Value = basetypes.NewInt64Null()
		return
	}

	v.Int64Value = basetypes.NewInt64PointerValue(s)
}

// * Int type

// SetInt sets the int64 value to the given int.
func (v *Int64Value) SetInt(s int) {
	v.Set(setInt64Value(s))
}

// SetIntPtr sets the int64 value to the given int pointer. If the pointer is nil, the value is set to null.
func (v *Int64Value) SetIntPtr(s *int) {
	if s == nil {
		v.Int64Value = basetypes.NewInt64Null()
		return
	}
	v.Int64Value = basetypes.NewInt64Value(setInt64Value(*s))
}

// GetInt returns converted int64 to int value.
func (v Int64Value) GetInt() int {
	return int(v.Get())
}

// GetIntPtr returns a converted int64 to int pointer. If the value is null or unknown, nil is returned.
func (v Int64Value) GetIntPtr() *int {
	if v.IsKnown() {
		i := int(v.Get())
		return &i
	}

	return nil
}

// * Int8 type

// SetInt8 sets the int64 value to the given int8.
func (v *Int64Value) SetInt8(s int8) {
	v.Set(setInt64Value(s))
}

// SetInt8Ptr sets the int64 value to the given int8 pointer. If the pointer is nil, the value is set to null.
func (v *Int64Value) SetInt8Ptr(s *int8) {
	if s == nil {
		v.Int64Value = basetypes.NewInt64Null()
		return
	}
	v.Int64Value = basetypes.NewInt64Value(setInt64Value(*s))
}

// GetInt8 returns converted int64 to int8 value.
func (v Int64Value) GetInt8() int8 {
	return int8(v.Get())
}

// GetInt8Ptr returns a converted int64 to int8 pointer. If the value is null or unknown, nil is returned.
func (v Int64Value) GetInt8Ptr() *int8 {
	if v.IsKnown() {
		i := int8(v.Get())
		return &i
	}

	return nil
}

// * Int16 type

// SetInt16 sets the int64 value to the given int16.
func (v *Int64Value) SetInt16(s int16) {
	v.Set(setInt64Value(s))
}

// SetInt16Ptr sets the int64 value to the given int16 pointer. If the pointer is nil, the value is set to null.
func (v *Int64Value) SetInt16Ptr(s *int16) {
	if s == nil {
		v.Int64Value = basetypes.NewInt64Null()
		return
	}
	v.Int64Value = basetypes.NewInt64Value(setInt64Value(*s))
}

// GetInt16 returns converted int64 to int16 value.
func (v Int64Value) GetInt16() int16 {
	return int16(v.Get())
}

// GetInt16Ptr returns a converted int64 to int16 pointer. If the value is null or unknown, nil is returned.
func (v Int64Value) GetInt16Ptr() *int16 {
	if v.IsKnown() {
		i := int16(v.Get())
		return &i
	}

	return nil
}

// * Int32 type

// SetInt32 sets the int64 value to the given int32.
func (v *Int64Value) SetInt32(s int32) {
	v.Set(setInt64Value(s))
}

// SetInt32Ptr sets the int64 value to the given int32 pointer. If the pointer is nil, the value is set to null.
func (v *Int64Value) SetInt32Ptr(s *int32) {
	if s == nil {
		v.Int64Value = basetypes.NewInt64Null()
		return
	}
	v.Int64Value = basetypes.NewInt64Value(setInt64Value(*s))
}

// GetInt32 returns converted int64 to int32 value.
func (v Int64Value) GetInt32() int32 {
	return int32(v.Get())
}

// GetInt32Ptr returns a converted int64 to int32 pointer. If the value is null or unknown, nil is returned.
func (v Int64Value) GetInt32Ptr() *int32 {
	if v.IsKnown() {
		i := int32(v.Get())
		return &i
	}

	return nil
}

// * Int64 type

// SetInt64 sets the int64 value to the given int64.
func (v *Int64Value) SetInt64(s int64) {
	v.Set(setInt64Value(s))
}

// SetInt64Ptr sets the int64 value to the given int64 pointer. If the pointer is nil, the value is set to null.
func (v *Int64Value) SetInt64Ptr(s *int64) {
	if s == nil {
		v.Int64Value = basetypes.NewInt64Null()
		return
	}
	v.Int64Value = basetypes.NewInt64Value(setInt64Value(*s))
}

// GetInt64 returns converted int64 to int64 value.
func (v Int64Value) GetInt64() int64 {
	return int64(v.Get())
}

// GetInt64Ptr returns a converted int64 to int64 pointer. If the value is null or unknown, nil is returned.
func (v Int64Value) GetInt64Ptr() *int64 {
	if v.IsKnown() {
		i := int64(v.Get())
		return &i
	}

	return nil
}

func setInt64Value[T intValues](s T) int64 {
	return int64(s)
}

// SetNull sets the Int64 value to null.
func (v *Int64Value) SetNull() {
	v.Int64Value = basetypes.NewInt64Null()
}

// SetUnknown sets the Int64 value to unknown.
func (v *Int64Value) SetUnknown() {
	v.Int64Value = basetypes.NewInt64Unknown()
}

// IsKnown returns true if the value is not null and not unknown.
func (v Int64Value) IsKnown() bool {
	return !v.Int64Value.IsNull() && !v.Int64Value.IsUnknown()
}
