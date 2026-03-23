// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// RaidControllerInventoryView RaidController
type RaidControllerInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ProductName string `json:"productName,omitempty"`
	SasAddress string `json:"sasAddress,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	AdapterNumber int `json:"adapterNumber,omitempty"`
	RaidPhysicalDrives []RaidPhysicalDriveInventoryView `json:"raidPhysicalDrives,omitempty"`
}

// RefreshLocalRaidEventView RefreshLocalRaidEvent
type RefreshLocalRaidEventView struct {
	Inventories []RaidControllerInventoryView `json:"inventories,omitempty"`
}

