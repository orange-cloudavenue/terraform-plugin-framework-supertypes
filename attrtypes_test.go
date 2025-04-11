/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package supertypes_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	supertypes "github.com/orange-cloudavenue/terraform-plugin-framework-supertypes"
)

func TestAttributeTypes(t *testing.T) {
	t.Parallel()

	type struct1 struct{}
	type struct2 struct {
		Name            types.String `tfsdk:"name"`
		ID              types.Int64  `tfsdk:"id"`
		IncludeProperty types.Bool   `tfsdk:"include_property"`
	}

	ctx := context.Background()
	got := supertypes.AttributeTypesMust[struct1](ctx)
	wanted := map[string]attr.Type{}

	if diff := cmp.Diff(got, wanted); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	_, err := supertypes.AttributeTypes[int](ctx)

	if err == nil {
		t.Fatalf("expected error")
	}

	got, err = supertypes.AttributeTypes[struct2](ctx)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	wanted = map[string]attr.Type{
		"name":             types.StringType,
		"id":               types.Int64Type,
		"include_property": types.BoolType,
	}

	if diff := cmp.Diff(got, wanted); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}
}

func TestElementType(t *testing.T) {
	t.Parallel()

	type String string
	type Int32 int32
	type Int64 int64
	type Bool bool
	type Float32 float32
	type Float64 float64
	type Invalid int

	ctx := context.Background()

	// TF framework type
	got, err := supertypes.ElementType[types.String](ctx)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if diff := cmp.Diff(got, types.StringType); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	// External custom type that uses basetypes.StringValue
	got, err = supertypes.ElementType[timetypes.RFC3339](ctx)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if diff := cmp.Diff(got, timetypes.RFC3339Type{}); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	// super type that uses basetypes.StringValue
	got, err = supertypes.ElementType[supertypes.StringValue](ctx)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if diff := cmp.Diff(got, supertypes.StringType{}); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	got, err = supertypes.ElementType[String](ctx)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if diff := cmp.Diff(got, supertypes.StringType{}); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	got, err = supertypes.ElementType[Int32](ctx)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if diff := cmp.Diff(got, supertypes.Int32Type{}); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	got, err = supertypes.ElementType[Int64](ctx)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if diff := cmp.Diff(got, supertypes.Int64Type{}); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	got, err = supertypes.ElementType[Bool](ctx)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if diff := cmp.Diff(got, supertypes.BoolType{}); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	got, err = supertypes.ElementType[Float32](ctx)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if diff := cmp.Diff(got, supertypes.Float32Type{}); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	got, err = supertypes.ElementType[Float64](ctx)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if diff := cmp.Diff(got, supertypes.Float64Type{}); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	_, err = supertypes.ElementType[Invalid](ctx)
	if err == nil {
		t.Fatalf("expected error")
	}
}
