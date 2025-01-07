/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package supertypes

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Must is a generic implementation of the Go Must idiom [1, 2]. It panics if
// the provided error is non-nil and returns x otherwise.
//
// [1]: https://pkg.go.dev/text/template#Must
// [2]: https://pkg.go.dev/regexp#MustCompile
func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

// MustDiag is a generic implementation of the Go Must idiom [1, 2]. It panics if
// the provided Diagnostics has errors and returns x otherwise.
//
// [1]: https://pkg.go.dev/text/template#Must
// [2]: https://pkg.go.dev/regexp#MustCompile
func MustDiag[T any](x T, diags diag.Diagnostics) T {
	return Must(x, DiagnosticsError(diags))
}

// DiagnosticsError returns an error containing all Diagnostic with SeverityError.
func DiagnosticsError(diags diag.Diagnostics) error {
	var errs []error

	for _, d := range diags.Errors() {
		errs = append(errs, errors.New(DiagnosticString(d)))
	}

	return errors.Join(errs...)
}

// DiagnosticString formats a Diagnostic
// If there is no `Detail`, only prints summary, otherwise prints both.
func DiagnosticString(d diag.Diagnostic) string {
	if d.Detail() == "" {
		return d.Summary()
	}
	return fmt.Sprintf("%s\n\n%s", d.Summary(), d.Detail())
}
