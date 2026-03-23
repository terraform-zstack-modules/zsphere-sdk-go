// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostNetworkInterfaceLldpInventoryView HostNetworkInterfaceLldp
type HostNetworkInterfaceLldpInventoryView struct {
	BaseInfoView
	BaseTimeView
	InterfaceUuid string `json:"interfaceUuid,omitempty"`
	Mode string `json:"mode,omitempty"`
	NeighborDevice HostNetworkInterfaceLldpRefInventoryView `json:"neighborDevice,omitempty"`
}

// QueryHostNetworkInterfaceLldpView QueryHostNetworkInterfaceLldp
type QueryHostNetworkInterfaceLldpView struct {
	Inventories []HostNetworkInterfaceLldpInventoryView `json:"inventories,omitempty"`
}

// GetHostNetworkInterfaceLldpView GetHostNetworkInterfaceLldp
type GetHostNetworkInterfaceLldpView struct {
	Lldp HostNetworkInterfaceLldpRefInventoryView `json:"lldp,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ChangeHostNetworkInterfaceLldpModeEventView ChangeHostNetworkInterfaceLldpModeEvent
type ChangeHostNetworkInterfaceLldpModeEventView struct {
	Inventories []HostNetworkInterfaceLldpInventoryView `json:"inventories,omitempty"`
}

