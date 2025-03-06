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
var _ basetypes.Int32Valuable = Int32Value{}

type Int32Value struct {
	basetypes.Int32Value
}

func (v Int32Value) Equal(o attr.Value) bool {
	other, ok := o.(Int32Value)

	if !ok {
		return false
	}

	return v.Int32Value.Equal(other.Int32Value)
}

func (v Int32Value) Type(ctx context.Context) attr.Type {
	return Int32Type{
		Int32Type: v.Int32Value.Type(ctx).(basetypes.Int32Type),
	}
}

func NewInt32Null() Int32Value {
	return Int32Value{
		Int32Value: basetypes.NewInt32Null(),
	}
}

func NewInt32Unknown() Int32Value {
	return Int32Value{
		Int32Value: basetypes.NewInt32Unknown(),
	}
}

func NewInt32Value(s int32) Int32Value {
	return Int32Value{
		Int32Value: basetypes.NewInt32Value(s),
	}
}

func NewInt32PointerValue(s *int32) Int32Value {
	return Int32Value{
		Int32Value: basetypes.NewInt32PointerValue(s),
	}
}

func NewInt32PointerValueOrNull(s *int32) Int32Value {
	if s == nil {
		return NewInt32Null()
	}

	return Int32Value{
		Int32Value: basetypes.NewInt32PointerValue(s),
	}
}

// * CustomFunc

// Get returns the known Int32 value.
func (v *Int32Value) Get() int32 {
	return v.Int32Value.ValueInt32()
}

// GetPtr returns a pointer to the known int64 value, nil for a
// null value, or a pointer to 0 for an unknown value.
func (v *Int32Value) GetPtr() *int32 {
	return v.Int32Value.ValueInt32Pointer()
}

// Set sets the Int32 value.
func (v *Int32Value) Set(s int32) {

	v.Int32Value = basetypes.NewInt32Value(s)
}

// SetPtr sets a pointer to the Int32 value.
func (v *Int32Value) SetPtr(s *int32) {
	if s == nil {
		v.Int32Value = basetypes.NewInt32Null()
		return
	}

	v.Int32Value = basetypes.NewInt32PointerValue(s)
}

// * Int type

// SetInt sets the int32 value to the given int.
func (v *Int32Value) SetInt(s int) {
	v.Set(setInt32Value(s))
}

// SetIntPtr sets the int32 value to the given int pointer. If the pointer is nil, the value is set to null.
func (v *Int32Value) SetIntPtr(s *int) {
	if s == nil {
		v.Int32Value = basetypes.NewInt32Null()
		return
	}
	v.Int32Value = basetypes.NewInt32Value(setInt32Value(*s))
}

// GetInt returns converted int32 to int value.
func (v Int32Value) GetInt() int {
	return int(v.Get())
}

// GetIntPtr returns a converted int32 to int pointer. If the value is null or unknown, nil is returned.
func (v Int32Value) GetIntPtr() *int {
	if v.IsKnown() {
		i := int(v.Get())
		return &i
	}

	return nil
}

// * Int8 type

// SetInt8 sets the int32 value to the given int8.
func (v *Int32Value) SetInt8(s int8) {
	v.Set(setInt32Value(s))
}

// SetInt8Ptr sets the int32 value to the given int8 pointer. If the pointer is nil, the value is set to null.
func (v *Int32Value) SetInt8Ptr(s *int8) {
	if s == nil {
		v.Int32Value = basetypes.NewInt32Null()
		return
	}
	v.Int32Value = basetypes.NewInt32Value(setInt32Value(*s))
}

// GetInt8 returns converted int32 to int8 value.
func (v Int32Value) GetInt8() int8 {
	return int8(v.Get())
}

// GetInt8Ptr returns a converted int32 to int8 pointer. If the value is null or unknown, nil is returned.
func (v Int32Value) GetInt8Ptr() *int8 {
	if v.IsKnown() {
		i := int8(v.Get())
		return &i
	}

	return nil
}

// * Int16 type

// SetInt16 sets the int32 value to the given int16.
func (v *Int32Value) SetInt16(s int16) {
	v.Set(setInt32Value(s))
}

// SetInt16Ptr sets the int32 value to the given int16 pointer. If the pointer is nil, the value is set to null.
func (v *Int32Value) SetInt16Ptr(s *int16) {
	if s == nil {
		v.Int32Value = basetypes.NewInt32Null()
		return
	}
	v.Int32Value = basetypes.NewInt32Value(setInt32Value(*s))
}

// GetInt16 returns converted int32 to int16 value.
func (v Int32Value) GetInt16() int16 {
	return int16(v.Get())
}

// GetInt16Ptr returns a converted int32 to int16 pointer. If the value is null or unknown, nil is returned.
func (v Int32Value) GetInt16Ptr() *int16 {
	if v.IsKnown() {
		i := int16(v.Get())
		return &i
	}

	return nil
}

// * Int32 type

// SetInt32 sets the int32 value to the given int32.
func (v *Int32Value) SetInt32(s int32) {
	v.Set(setInt32Value(s))
}

// SetInt32Ptr sets the int32 value to the given int32 pointer. If the pointer is nil, the value is set to null.
func (v *Int32Value) SetInt32Ptr(s *int32) {
	if s == nil {
		v.Int32Value = basetypes.NewInt32Null()
		return
	}
	v.Int32Value = basetypes.NewInt32Value(setInt32Value(*s))
}

// GetInt32 returns converted int32 to int32 value.
func (v Int32Value) GetInt32() int32 {
	return int32(v.Get())
}

// GetInt32Ptr returns a converted int32 to int32 pointer. If the value is null or unknown, nil is returned.
func (v Int32Value) GetInt32Ptr() *int32 {
	if v.IsKnown() {
		i := int32(v.Get())
		return &i
	}

	return nil
}

// * Int64 type

// SetInt64 sets the int32 value to the given int64.
func (v *Int32Value) SetInt64(s int64) {
	v.Set(setInt32Value(s))
}

// SetInt64Ptr sets the int32 value to the given int64 pointer. If the pointer is nil, the value is set to null.
func (v *Int32Value) SetInt64Ptr(s *int64) {
	if s == nil {
		v.Int32Value = basetypes.NewInt32Null()
		return
	}
	v.Int32Value = basetypes.NewInt32Value(setInt32Value(*s))
}

// GetInt64 returns converted int32 to int64 value.
func (v Int32Value) GetInt64() int64 {
	return int64(v.Get())
}

// GetInt64Ptr returns a converted int32 to int64 pointer. If the value is null or unknown, nil is returned.
func (v Int32Value) GetInt64Ptr() *int64 {
	if v.IsKnown() {
		i := int64(v.Get())
		return &i
	}

	return nil
}

func setInt32Value[T intValues](s T) int32 {
	return int32(s)
}

// SetNull sets the Int32 value to null.
func (v *Int32Value) SetNull() {
	v.Int32Value = basetypes.NewInt32Null()
}

// SetUnknown sets the Int32 value to unknown.
func (v *Int32Value) SetUnknown() {
	v.Int32Value = basetypes.NewInt32Unknown()
}

// IsKnown returns true if the value is not null and not unknown.
func (v Int32Value) IsKnown() bool {
	return !v.Int32Value.IsNull() && !v.Int32Value.IsUnknown()
}
