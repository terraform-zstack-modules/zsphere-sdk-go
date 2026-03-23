// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BaremetalBondingInventoryView BaremetalBonding
type BaremetalBondingInventoryView struct {
	BaseInfoView
	BaseTimeView
	ChassisUuid string `json:"chassisUuid,omitempty"`
	Mode int `json:"mode,omitempty"`
	Slaves string `json:"slaves,omitempty"`
	Opts string `json:"opts,omitempty"`
}

// CreateBaremetalBondingEventView CreateBaremetalBondingEvent
type CreateBaremetalBondingEventView struct {
	Inventory BaremetalBondingInventoryView `json:"inventory,omitempty"`
}

// QueryBaremetalBondingView QueryBaremetalBonding
type QueryBaremetalBondingView struct {
	Inventories []BaremetalBondingInventoryView `json:"inventories,omitempty"`
}

