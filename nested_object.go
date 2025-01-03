/*
 * SPDX-FileCopyrightText: Copyright (c) Orange Business Services SA
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at <https://www.mozilla.org/en-US/MPL/2.0/>
 * or see the "LICENSE" file for more details.
 */

package supertypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// NestedObjectType extends the Type interface for types that represent nested Objects.
type NestedObjectType interface {
	attr.Type

	// NewObjectPtr returns a new, empty value as an object pointer (Go *struct).
	NewObjectPtr(context.Context) (any, diag.Diagnostics)

	// NewObjectSlice returns a new value as an object slice (Go []*struct).
	NewObjectSlice(context.Context, int, int) (any, diag.Diagnostics)

	// NullValue returns a Null Value.
	NullValue(context.Context) (attr.Value, diag.Diagnostics)

	// ValueFromObjectPtr returns a Value given an object pointer (Go *struct).
	ValueFromObjectPtr(context.Context, any) (attr.Value, diag.Diagnostics)

	// ValueFromObjectSlice returns a Value given an object pointer (Go []*struct).
	ValueFromObjectSlice(context.Context, any) (attr.Value, diag.Diagnostics)
}

// NestedObjectValue extends the Value interface for values that represent nested Objects.
type NestedObjectValue interface {
	attr.Value
}

// valueWithElements extends the Value interface for values that have an Elements method.
type valueWithElements interface {
	attr.Value

	Elements() []attr.Value
}

// valueWithElementsMap extends the Value interface for values that have an ElementsMap method.
type valueWithElementsMap interface {
	attr.Value

	Elements() map[string]attr.Value
}
