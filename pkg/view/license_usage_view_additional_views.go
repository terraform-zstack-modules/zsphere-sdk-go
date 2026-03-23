// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LicenseUsageViewView LicenseUsageView
type LicenseUsageViewView struct {
	QuotaType string `json:"quotaType,omitempty"`
	Quota int64 `json:"quota,omitempty"`
	Used int64 `json:"used,omitempty"`
	Available int64 `json:"available,omitempty"`
}

