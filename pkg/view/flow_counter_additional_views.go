// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// FlowCounterView FlowCounter
type FlowCounterView struct {
	Device string `json:"device,omitempty"`
	TotalEntries string `json:"totalEntries,omitempty"`
	TotalPkts string `json:"totalPkts,omitempty"`
	TotalBytes string `json:"totalBytes,omitempty"`
}

