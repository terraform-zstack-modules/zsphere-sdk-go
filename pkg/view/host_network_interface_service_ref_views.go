// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostNetworkInterfaceServiceRefInventoryView HostNetworkInterfaceServiceRef
type HostNetworkInterfaceServiceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	InterfaceUuid string `json:"interfaceUuid,omitempty"`
	VlanId int `json:"vlanId,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
}

// SetServiceTypeOnHostNetworkInterfaceEventView SetServiceTypeOnHostNetworkInterfaceEvent
type SetServiceTypeOnHostNetworkInterfaceEventView struct {
	Inventory []HostNetworkInterfaceServiceRefInventoryView `json:"inventory,omitempty"`
}

