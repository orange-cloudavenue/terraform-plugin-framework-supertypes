## 0.5.0 (Unreleased)
### :rotating_light: **Breaking Changes**

* `go` - Bump golang version from `1.21.0` to `1.22.0` (GH-7)
### :information_source: **Notes**

* `core` - The `FrangipaneTeam/terraform-plugin-framework-supertypes`  has been renamed to `orange-cloudavenue/terraform-plugin-framework-supertypes`. The plugin has been moved to be under the `orange-cloudavenue` organization. (GH-5)

### :dependabot: **Dependencies**

* deps: bumps github.com/hashicorp/terraform-plugin-go from 0.23.0 to 0.25.0 (GH-3)

## 0.4.0 (August 20, 2024)
### :rotating_light: **Breaking Changes**

* `golang` - Upgrade to Golang 1.21. (GH-71)
### :information_source: **Notes**

* `linters` - Remove deprecated linters. (GH-71)

### :dependabot: **Dependencies**

* deps: bumps actions/checkout from 4.1.1 to 4.1.7 (GH-68)
* deps: bumps actions/download-artifact from 4.1.1 to 4.1.8 (GH-67)
* deps: bumps dependabot/fetch-metadata from 1.6.0 to 2.2.0 (GH-66)
* deps: bumps github.com/hashicorp/terraform-plugin-go from 0.21.0 to 0.23.0 (GH-69)
* deps: bumps github.com/stretchr/testify from 1.8.4 to 1.9.0 (GH-60)
* deps: bumps golangci/golangci-lint-action from 3 to 6 (GH-70)
* deps: bumps ncipollo/release-action from 1.13.0 to 1.14.0 (GH-57)

## 0.3.1 (February  2, 2024)

### :bug: **Bug Fixes**

* `MapNestedValueOF` - Fix wrong parameter for the funcs `Get()` and `Set()`. (GH-55)

### :dependabot: **Dependencies**

* deps: bumps github.com/hashicorp/terraform-plugin-go from 0.20.0 to 0.21.0 (GH-54)

## 0.3.0 (January 17, 2024)

### :rocket: **New Features**

* `attribute/ListAttributeOf` - Add list attribute with generic type for automatically cast value to type of attribute. (GH-49)
* `attribute/MapAttributeOf` - Add map attribute with generic type for automatically cast value to type of attribute. (GH-49)
* `attribute/SetAttributeOf` - Add set attribute with generic type for automatically cast value to type of attribute. (GH-49)

### :dependabot: **Dependencies**

* deps: bumps actions/checkout from 4.1.0 to 4.1.1 (GH-39)
* deps: bumps actions/download-artifact from 3.0.2 to 4.1.1 (GH-51)
* deps: bumps actions/setup-go from 4 to 5 (GH-45)
* deps: bumps actions/upload-artifact from 3 to 4 (GH-47)
* deps: bumps github.com/fatih/color from 1.15.0 to 1.16.0 (GH-42)
* deps: bumps github.com/hashicorp/terraform-plugin-framework from 1.4.2 to 1.5.0 (GH-50)
* deps: bumps github.com/hashicorp/terraform-plugin-go from 0.19.0 to 0.19.1 (GH-44)

## 0.2.0 (October 31, 2023)
## 0.1.1 (October 23, 2023)

### :dependabot: **Dependencies**

* deps: bumps actions/checkout from 3 to 4 ([GH-29](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/29))
* deps: bumps actions/checkout from 4.0.0 to 4.1.0 ([GH-32](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/32))
* deps: bumps github.com/google/go-cmp from 0.5.9 to 0.6.0 ([GH-35](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/35))
* deps: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.5 to 1.4.0 ([GH-31](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/31))
* deps: bumps github.com/hashicorp/terraform-plugin-framework from 1.4.0 to 1.4.1 ([GH-33](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/33))
* deps: bumps github.com/stretchr/testify from 1.6.1 to 1.8.4 ([GH-28](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/28))
* deps: bumps ncipollo/release-action from 1.12.0 to 1.13.0 ([GH-27](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/27))
* deps: bumps stefanzweifel/git-auto-commit-action from 4 to 5 ([GH-34](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/34))

## 0.1.0 (August 22, 2023)
### :rotating_light: **Breaking Changes**

* `attribute/string` - Now `Set` func set attribute to null if value is empty string ([GH-26](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/26))

### :rocket: **New Features**

* `attribute/bool` - Add func `SetPtr` ([GH-26](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/26))
* `attribute/float64` - Add func `SetFloat32`, `SetFloat64`, `SetFloat32Ptr`, `SetFloat64Ptr`, `GetFloat32`, `GetFloat64`, `GetFloat32Ptr`, `GetFloat64Ptr` ([GH-26](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/26))
* `attribute/int64` - Add funcs `SetPtr`, `SetInt`, `SetInt8`, `SetInt16`, `SetInt32`, `SetInt64`, `SetIntPtr`, `SetInt8Ptr`, `SetInt16Ptr`, `SetInt32Ptr`, `SetInt64Ptr`, `GetInt`, `GetInt8`, `GetInt16`, `GetInt32`, `GetInt64`, `GetIntPtr`, `GetInt8Ptr`, `GetInt16Ptr`, `GetInt32Ptr`, `GetInt64Ptr` ([GH-26](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/26))
* `attribute/string` - Add func `SetPtr` ([GH-26](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/26))

### :dependabot: **Dependencies**

* deps: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.2 to 1.3.3 ([GH-21](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/21))
* deps: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.3 to 1.3.4 ([GH-24](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/24))
* deps: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.4 to 1.3.5 ([GH-25](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/25))

## 0.0.5 (July 19, 2023)

## 0.0.4
