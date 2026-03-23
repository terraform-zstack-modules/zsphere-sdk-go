// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostPhysicalMemoryInventoryView HostPhysicalMemory
type HostPhysicalMemoryInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostUuid string `json:"hostUuid,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Size string `json:"size,omitempty"`
	Speed string `json:"speed,omitempty"`
	ClockSpeed string `json:"clockSpeed,omitempty"`
	Locator string `json:"locator,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
	Rank string `json:"rank,omitempty"`
	Voltage string `json:"voltage,omitempty"`
	Type string `json:"type,omitempty"`
}

// QueryHostPhysicalMemoryView QueryHostPhysicalMemory
type QueryHostPhysicalMemoryView struct {
	Inventories []HostPhysicalMemoryInventoryView `json:"inventories,omitempty"`
}

// GetHostPhysicalMemoryFactsView GetHostPhysicalMemoryFacts
type GetHostPhysicalMemoryFactsView struct {
	Inventories []HostPhysicalMemoryInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

