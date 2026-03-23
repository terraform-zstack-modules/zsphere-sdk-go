// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SharedBlockCapacityInventoryView SharedBlockCapacity
type SharedBlockCapacityInventoryView struct {
	BaseInfoView
	BaseTimeView
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
}

