// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EipInventoryView Eip
type EipInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	State string `json:"state,omitempty"`
	VipIp string `json:"vipIp,omitempty"`
	GuestIp string `json:"guestIp,omitempty"`
}

// ChangeEipStateEventView ChangeEipStateEvent
type ChangeEipStateEventView struct {
	Inventory EipInventoryView `json:"inventory,omitempty"`
}

// CreateEipEventView CreateEipEvent
type CreateEipEventView struct {
	Inventory EipInventoryView `json:"inventory,omitempty"`
}

// GetVpcAttachedEipView GetVpcAttachedEip
type GetVpcAttachedEipView struct {
	Inventories []EipInventoryView `json:"inventories,omitempty"`
}

// AttachEipEventView AttachEipEvent
type AttachEipEventView struct {
	Inventory EipInventoryView `json:"inventory,omitempty"`
}

// UpdateEipEventView UpdateEipEvent
type UpdateEipEventView struct {
	Inventory EipInventoryView `json:"inventory,omitempty"`
}

// DeleteEipEventView DeleteEipEvent
type DeleteEipEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmNicAttachableEipsView GetVmNicAttachableEips
type GetVmNicAttachableEipsView struct {
	Inventories []EipInventoryView `json:"inventories,omitempty"`
}

// QueryEipView QueryEip
type QueryEipView struct {
	Inventories []EipInventoryView `json:"inventories,omitempty"`
}

// DetachEipEventView DetachEipEvent
type DetachEipEventView struct {
	Inventory EipInventoryView `json:"inventory,omitempty"`
}

