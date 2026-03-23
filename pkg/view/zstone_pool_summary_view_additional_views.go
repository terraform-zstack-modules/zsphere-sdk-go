// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZStonePoolSummaryViewView ZStonePoolSummaryView
type ZStonePoolSummaryViewView struct {
	Count int `json:"count,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	Inventories []interface{} `json:"inventories,omitempty"`
}

