// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZceXInventoryView ZceX
type ZceXInventoryView struct {
	BaseInfoView
	BaseTimeView
	ManagementIp string `json:"managementIp,omitempty"`
	ApiPort int `json:"apiPort,omitempty"`
}

// RemoveZceXEventView RemoveZceXEvent
type RemoveZceXEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddZceXEventView AddZceXEvent
type AddZceXEventView struct {
	Inventory ZceXInventoryView `json:"inventory,omitempty"`
}

// QueryZceXView QueryZceX
type QueryZceXView struct {
	Inventories []ZceXInventoryView `json:"inventories,omitempty"`
}

