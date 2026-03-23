// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VniRangeInventoryView VniRange
type VniRangeInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	StartVni int `json:"startVni,omitempty"`
	EndVni int `json:"endVni,omitempty"`
	L2NetworkUuid string `json:"l2NetworkUuid,omitempty"`
}

// UpdateVniRangeEventView UpdateVniRangeEvent
type UpdateVniRangeEventView struct {
	Inventory VniRangeInventoryView `json:"inventory,omitempty"`
}

// QueryVniRangeView QueryVniRange
type QueryVniRangeView struct {
	Inventories []VniRangeInventoryView `json:"inventories,omitempty"`
}

// CreateVniRangeEventView CreateVniRangeEvent
type CreateVniRangeEventView struct {
	Inventory VniRangeInventoryView `json:"inventory,omitempty"`
}

// DeleteVniRangeEventView DeleteVniRangeEvent
type DeleteVniRangeEventView struct {
	Success bool `json:"success,omitempty"`
}

