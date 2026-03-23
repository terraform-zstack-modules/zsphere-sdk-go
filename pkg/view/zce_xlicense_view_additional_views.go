// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZceXLicenseViewView ZceXLicenseView
type ZceXLicenseViewView struct {
	Platform ZceXPlatformLicenseViewView `json:"platform,omitempty"`
}

