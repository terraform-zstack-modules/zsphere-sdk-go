// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BaremetalNicInventoryView BaremetalNic
type BaremetalNicInventoryView struct {
	BaseInfoView
	BaseTimeView
	BaremetalInstanceUuid string `json:"baremetalInstanceUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	BaremetalBondingUuid string `json:"baremetalBondingUuid,omitempty"`
	Mac string `json:"mac,omitempty"`
	Ip string `json:"ip,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	Pxe bool `json:"pxe,omitempty"`
}

