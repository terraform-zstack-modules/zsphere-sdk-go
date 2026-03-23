// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VtepInventoryView Vtep
type VtepInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostUuid string `json:"hostUuid,omitempty"`
	VtepIp string `json:"vtepIp,omitempty"`
	Port int `json:"port,omitempty"`
	Type string `json:"type,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	PoolUuid string `json:"poolUuid,omitempty"`
}

// QueryVtepView QueryVtep
type QueryVtepView struct {
	Inventories []VtepInventoryView `json:"inventories,omitempty"`
}

// CreateVxlanVtepEventView CreateVxlanVtepEvent
type CreateVxlanVtepEventView struct {
	Inventory VtepInventoryView `json:"inventory,omitempty"`
}

