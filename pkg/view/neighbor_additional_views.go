// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NeighborView Neighbor
type NeighborView struct {
	Id string `json:"id,omitempty"`
	Priority string `json:"priority,omitempty"`
	State string `json:"state,omitempty"`
	DeadTime string `json:"deadTime,omitempty"`
	NeighborAddress string `json:"neighborAddress,omitempty"`
	Device string `json:"device,omitempty"`
}

