// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// UpdateLicenseViewView UpdateLicenseView
type UpdateLicenseViewView struct {
	License string `json:"license,omitempty"`
	HandleBy string `json:"handleBy,omitempty"`
}

