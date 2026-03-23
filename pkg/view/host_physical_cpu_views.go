// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostPhysicalCpuInventoryView HostPhysicalCpu
type HostPhysicalCpuInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostUuid string `json:"hostUuid,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
	SocketDesignation string `json:"socketDesignation,omitempty"`
	Version string `json:"version,omitempty"`
	CurrentSpeed string `json:"currentSpeed,omitempty"`
	CoreCount int `json:"coreCount,omitempty"`
	ThreadCount int `json:"threadCount,omitempty"`
}

// QueryHostPhysicalCpuView QueryHostPhysicalCpu
type QueryHostPhysicalCpuView struct {
	Inventories []HostPhysicalCpuInventoryView `json:"inventories,omitempty"`
}

