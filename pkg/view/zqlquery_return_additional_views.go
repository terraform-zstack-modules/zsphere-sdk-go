// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZQLQueryReturnView ZQLQueryReturn
type ZQLQueryReturnView struct {
	Inventories ListView `json:"inventories,omitempty"`
	InventoryCounts map[string]int64 `json:"inventoryCounts,omitempty"`
	Total int64 `json:"total,omitempty"`
	ReturnWith interface{} `json:"returnWith,omitempty"`
	Name string `json:"name,omitempty"`
}

