// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostNetworkBondingInventoryView HostNetworkBonding
type HostNetworkBondingInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostUuid string `json:"hostUuid,omitempty"`
	BondingName string `json:"bondingName,omitempty"`
	BondingType string `json:"bondingType,omitempty"`
	Speed int64 `json:"speed,omitempty"`
	Mode string `json:"mode,omitempty"`
	XmitHashPolicy string `json:"xmitHashPolicy,omitempty"`
	MiiStatus string `json:"miiStatus,omitempty"`
	Mac string `json:"mac,omitempty"`
	IpAddresses []string `json:"ipAddresses,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	CallBackIp string `json:"callBackIp,omitempty"`
	Miimon int64 `json:"miimon,omitempty"`
	Type string `json:"type,omitempty"`
	AllSlavesActive bool `json:"allSlavesActive,omitempty"`
	Description string `json:"description,omitempty"`
	Slaves []HostNetworkInterfaceInventoryView `json:"slaves,omitempty"`
}

// AttachNicToBondingEventView AttachNicToBondingEvent
type AttachNicToBondingEventView struct {
	Inventory HostNetworkBondingInventoryView `json:"inventory,omitempty"`
}

// DetachNicFromBondingEventView DetachNicFromBondingEvent
type DetachNicFromBondingEventView struct {
	Inventory HostNetworkBondingInventoryView `json:"inventory,omitempty"`
}

// UpdateBondingEventView UpdateBondingEvent
type UpdateBondingEventView struct {
	Inventory HostNetworkBondingInventoryView `json:"inventory,omitempty"`
}

// QueryHostNetworkBondingView QueryHostNetworkBonding
type QueryHostNetworkBondingView struct {
	Inventories []HostNetworkBondingInventoryView `json:"inventories,omitempty"`
}

// UpdateVirtualSwitchUplinkBondingsEventView UpdateVirtualSwitchUplinkBondingsEvent
type UpdateVirtualSwitchUplinkBondingsEventView struct {
	Inventories []HostNetworkBondingInventoryView `json:"inventories,omitempty"`
}

// SetIpOnHostNetworkBondingEventView SetIpOnHostNetworkBondingEvent
type SetIpOnHostNetworkBondingEventView struct {
	Inventory HostNetworkBondingInventoryView `json:"inventory,omitempty"`
}

// CreateBondingEventView CreateBondingEvent
type CreateBondingEventView struct {
	Inventory []HostNetworkBondingInventoryView `json:"inventory,omitempty"`
}

// GetCandidateNetworkBondingsView GetCandidateNetworkBondings
type GetCandidateNetworkBondingsView struct {
	Inventories []HostNetworkBondingInventoryView `json:"inventories,omitempty"`
}

