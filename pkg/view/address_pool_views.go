// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AddressPoolInventoryView AddressPool
type AddressPoolInventoryView struct {
	BaseInfoView
	BaseTimeView
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Description string `json:"description,omitempty"`
	StartIp string `json:"startIp,omitempty"`
	EndIp string `json:"endIp,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	NetworkCidr string `json:"networkCidr,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	AddressMode string `json:"addressMode,omitempty"`
	PrefixLen int `json:"prefixLen,omitempty"`
	IpRangeType string `json:"ipRangeType,omitempty"`
}

// QueryAddressPoolView QueryAddressPool
type QueryAddressPoolView struct {
	Inventories []AddressPoolInventoryView `json:"inventories,omitempty"`
}

