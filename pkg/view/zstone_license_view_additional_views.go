// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZStoneLicenseViewView ZStoneLicenseView
type ZStoneLicenseViewView struct {
	Platform ZStoneLicenseInventoryView `json:"platform,omitempty"`
	AddOns []ZStoneLicenseInventoryView `json:"addOns,omitempty"`
}

