// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZceXPoolSummaryViewView ZceXPoolSummaryView
type ZceXPoolSummaryViewView struct {
	Count int `json:"count,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	Inventories []interface{} `json:"inventories,omitempty"`
}

