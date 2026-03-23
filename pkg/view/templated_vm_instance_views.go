// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TemplatedVmInstanceInventoryView TemplatedVmInstance
type TemplatedVmInstanceInventoryView struct {
	BaseInfoView
	BaseTimeView
	ZoneUuid string `json:"zoneUuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

// UpdateTemplatedVmInstanceEventView UpdateTemplatedVmInstanceEvent
type UpdateTemplatedVmInstanceEventView struct {
	Inventory TemplatedVmInstanceInventoryView `json:"inventory,omitempty"`
}

// DeleteTemplatedVmInstanceEventView DeleteTemplatedVmInstanceEvent
type DeleteTemplatedVmInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryTemplatedVmInstanceView QueryTemplatedVmInstance
type QueryTemplatedVmInstanceView struct {
	Inventories []TemplatedVmInstanceInventoryView `json:"inventories,omitempty"`
}

