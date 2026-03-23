// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZBoxInventoryView ZBox
type ZBoxInventoryView struct {
	BaseInfoView
	BaseTimeView
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	LocationRefs []ZBoxLocationRefInventoryView `json:"locationRefs,omitempty"`
	MountPath string `json:"mountPath,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
}

// AddZBoxEventView AddZBoxEvent
type AddZBoxEventView struct {
	Inventory ZBoxInventoryView `json:"inventory,omitempty"`
}

// QueryZBoxView QueryZBox
type QueryZBoxView struct {
	Inventories []ZBoxInventoryView `json:"inventories,omitempty"`
}

// SyncZBoxCapacityEventView SyncZBoxCapacityEvent
type SyncZBoxCapacityEventView struct {
	Inventory ZBoxInventoryView `json:"inventory,omitempty"`
}

