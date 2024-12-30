package supertypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ attr.Type = missingType{}

// missingType is a placeholder attr.Type implementation for when type
// information is missing. This type is never valid for real usage, which is why
// it is unexported, but it is primarily used by other base types when an
// expected attr.Type field is nil for panic prevention and troubleshooting.
// Ideally those other base type implementations would make it impossible to
// create a situation which needs this, but those exported APIs are protected by
// compatibility promises until a major version.
type missingType struct{}

// ApplyTerraform5AttributePathStep always returns an error.
func (t missingType) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	return nil, fmt.Errorf("cannot apply AttributePathStep %T to %s", step, t.String())
}

// Equal returns true if the given type is equivalent.
func (t missingType) Equal(o attr.Type) bool {
	_, ok := o.(missingType)

	return ok
}

// String returns a human readable string of the type.
func (t missingType) String() string {
	return "!!! SUPER MISSING TYPE !!!"
}

// TerraformType returns DynamicPseudoType.
func (t missingType) TerraformType(_ context.Context) tftypes.Type {
	// Ideally, upstream would implement an "invalid" primitive type for this
	// situation, but DynamicPseudoType is an alternative unexpected type in
	// the framework until it potentially implements its own dynamic type
	// handling.
	return tftypes.DynamicPseudoType
}

// ValueFromTerraform always returns an error.
func (t missingType) ValueFromTerraform(_ context.Context, _ tftypes.Value) (attr.Value, error) {
	return missingValue{}, fmt.Errorf("missing type information; cannot create value")
}

// ValueType returns the missingValue type.
func (t missingType) ValueType(_ context.Context) attr.Value {
	return missingValue{}
}

var _ attr.Value = missingValue{}

// missingValue is a placeholder attr.Value implementation for when type
// information is missing. This type is never valid for real usage, which is why
// it is unexported, but it is primarily used by other base types when an
// expected attr.Value field is nil for panic prevention and troubleshooting.
// Ideally those other base type implementations would make it impossible to
// create a situation which needs this, but those exported APIs are protected by
// compatibility promises until a major version.
type missingValue struct{}

// Equal returns true if the given value is a missingValue.
func (v missingValue) Equal(o attr.Value) bool {
	_, ok := o.(missingValue)

	return ok
}

// IsNull returns false.
func (v missingValue) IsNull() bool {
	// Short of causing a panic, this method must choose a return value and
	// false was chosen so it is always "known".
	return false
}

// IsUnknown returns false.
func (v missingValue) IsUnknown() bool {
	// Short of causing a panic, this method must choose a return value and
	// false was chosen so it is always "known".
	return false
}

// String returns a human-readable representation of the value.
//
// The string returned here is not protected by any compatibility guarantees,
// and is intended for logging and error reporting.
func (v missingValue) String() string {
	return "!!! MISSING VALUE !!!"
}

// ToTerraformValue always returns an error.
func (v missingValue) ToTerraformValue(_ context.Context) (tftypes.Value, error) {
	return tftypes.Value{}, nil
}

// Type returns missingType.
func (v missingValue) Type(_ context.Context) attr.Type {
	return missingType{}
}
