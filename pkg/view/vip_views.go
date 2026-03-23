// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VipInventoryView Vip
type VipInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Ip string `json:"ip,omitempty"`
	State string `json:"state,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	PrefixLen int `json:"prefixLen,omitempty"`
	ServiceProvider string `json:"serviceProvider,omitempty"`
	PeerL3NetworkUuids []string `json:"peerL3NetworkUuids,omitempty"`
	ServicesRefs []VipNetworkServicesRefInventoryView `json:"servicesRefs,omitempty"`
	UseFor string `json:"useFor,omitempty"`
	System bool `json:"system,omitempty"`
}

// UpdateVipEventView UpdateVipEvent
type UpdateVipEventView struct {
	Inventory VipInventoryView `json:"inventory,omitempty"`
}

// CreateVipEventView CreateVipEvent
type CreateVipEventView struct {
	Inventory VipInventoryView `json:"inventory,omitempty"`
}

// ChangeVipStateEventView ChangeVipStateEvent
type ChangeVipStateEventView struct {
	Inventory VipInventoryView `json:"inventory,omitempty"`
}

// GetVpcAttachedVipView GetVpcAttachedVip
type GetVpcAttachedVipView struct {
	Inventories []VipInventoryView `json:"inventories,omitempty"`
}

// QueryVipView QueryVip
type QueryVipView struct {
	Inventories []VipInventoryView `json:"inventories,omitempty"`
}

// DeleteVipEventView DeleteVipEvent
type DeleteVipEventView struct {
	Success bool `json:"success,omitempty"`
}

