// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// QuotaUsageView QuotaUsage
type QuotaUsageView struct {
	Name string `json:"name,omitempty"`
	Total int64 `json:"total,omitempty"`
	Used int64 `json:"used,omitempty"`
}

