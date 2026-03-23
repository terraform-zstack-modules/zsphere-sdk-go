// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostKernelInterfaceInventoryView HostKernelInterface
type HostKernelInterfaceInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	L2NetworkUuid string `json:"l2NetworkUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	UsedIps []HostKernelInterfaceUsedIpInventoryView `json:"usedIps,omitempty"`
	TrafficTypes []string `json:"trafficTypes,omitempty"`
}

// DeleteHostKernelInterfaceEventView DeleteHostKernelInterfaceEvent
type DeleteHostKernelInterfaceEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateHostKernelInterfaceEventView CreateHostKernelInterfaceEvent
type CreateHostKernelInterfaceEventView struct {
	Inventory HostKernelInterfaceInventoryView `json:"inventory,omitempty"`
}

// QueryHostKernelInterfaceView QueryHostKernelInterface
type QueryHostKernelInterfaceView struct {
	Inventories []HostKernelInterfaceInventoryView `json:"inventories,omitempty"`
}

// UpdateHostKernelInterfaceEventView UpdateHostKernelInterfaceEvent
type UpdateHostKernelInterfaceEventView struct {
	Inventory HostKernelInterfaceInventoryView `json:"inventory,omitempty"`
}

