// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// UplinkGroupInventoryView UplinkGroup
type UplinkGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	InterfaceName string `json:"interfaceName,omitempty"`
	Type string `json:"type,omitempty"`
	BondingUuid string `json:"bondingUuid,omitempty"`
	InterfaceUuid string `json:"interfaceUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	L2NetworkUuid string `json:"l2NetworkUuid,omitempty"`
	L2ProviderType string `json:"l2ProviderType,omitempty"`
	BridgeName string `json:"bridgeName,omitempty"`
	SkipDeletion bool `json:"skipDeletion,omitempty"`
}

// QueryUplinkGroupView QueryUplinkGroup
type QueryUplinkGroupView struct {
	Inventories []UplinkGroupInventoryView `json:"inventories,omitempty"`
}

// UpdateVirtualSwitchUplinkGroupEventView UpdateVirtualSwitchUplinkGroupEvent
type UpdateVirtualSwitchUplinkGroupEventView struct {
	Inventory UplinkGroupInventoryView `json:"inventory,omitempty"`
}

