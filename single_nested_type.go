// -------------------------------------------------------------------- //
// !      DO NOT EDIT. This file is auto-generated from template      ! //
// -------------------------------------------------------------------- //
package supertypes

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.ObjectTypable = SingleNestedType{}

type SingleNestedType struct {
	basetypes.ObjectType
}

func (t SingleNestedType) Equal(o attr.Type) bool {
	other, ok := o.(basetypes.ObjectType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other)
}

func (t SingleNestedType) String() string {
	var res strings.Builder
	res.WriteString("types.ObjectType[")
	keys := make([]string, 0, len(t.AttrTypes))
	for k := range t.AttrTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for pos, key := range keys {
		if pos != 0 {
			res.WriteString(", ")
		}
		res.WriteString(`"` + key + `":`)
		res.WriteString(t.AttrTypes[key].String())
	}
	res.WriteString("]")
	return res.String()
}

func (t SingleNestedType) ValueFromObject(_ context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	// CustomSingleValue defined in the value type section
	value := SingleNestedValue{
		ObjectValue: in,
	}

	return value, nil
}

func (t SingleNestedType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ObjectType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	ObjectValue, ok := attrValue.(basetypes.ObjectValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	value, diags := t.ValueFromObject(ctx, ObjectValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting SingleValue to SingleValuable: %v", diags)
	}

	return value, nil
}