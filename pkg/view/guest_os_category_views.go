// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// GuestOsCategoryInventoryView GuestOsCategory
type GuestOsCategoryInventoryView struct {
	BaseInfoView
	BaseTimeView
	Platform string `json:"platform,omitempty"`
	OsRelease string `json:"osRelease,omitempty"`
	Version string `json:"version,omitempty"`
}

