// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostNetworkBondingServiceRefInventoryView HostNetworkBondingServiceRef
type HostNetworkBondingServiceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	BondingUuid string `json:"bondingUuid,omitempty"`
	VlanId int `json:"vlanId,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
}

// SetServiceTypeOnHostNetworkBondingEventView SetServiceTypeOnHostNetworkBondingEvent
type SetServiceTypeOnHostNetworkBondingEventView struct {
	Inventory []HostNetworkBondingServiceRefInventoryView `json:"inventory,omitempty"`
}

