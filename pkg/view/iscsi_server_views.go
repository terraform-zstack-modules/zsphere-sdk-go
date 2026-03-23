// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// IscsiServerInventoryView IscsiServer
type IscsiServerInventoryView struct {
	BaseInfoView
	BaseTimeView
	Ip string `json:"ip,omitempty"`
	Port int `json:"port,omitempty"`
	ChapUserName string `json:"chapUserName,omitempty"`
	ChapUserPassword string `json:"chapUserPassword,omitempty"`
	State string `json:"state,omitempty"`
	IscsiTargets []IscsiTargetInventoryView `json:"iscsiTargets,omitempty"`
	IscsiClusterRefs []IscsiServerClusterRefInventoryView `json:"iscsiClusterRefs,omitempty"`
}

// RefreshIscsiServerEventView RefreshIscsiServerEvent
type RefreshIscsiServerEventView struct {
	Inventory IscsiServerInventoryView `json:"inventory,omitempty"`
}

// AddIscsiServerEventView AddIscsiServerEvent
type AddIscsiServerEventView struct {
	Inventory IscsiServerInventoryView `json:"inventory,omitempty"`
}

// AttachIscsiServerToClusterEventView AttachIscsiServerToClusterEvent
type AttachIscsiServerToClusterEventView struct {
	Inventory IscsiServerInventoryView `json:"inventory,omitempty"`
}

// QueryIscsiServerView QueryIscsiServer
type QueryIscsiServerView struct {
	Inventories []IscsiServerInventoryView `json:"inventories,omitempty"`
}

// UpdateIscsiServerEventView UpdateIscsiServerEvent
type UpdateIscsiServerEventView struct {
	Inventory IscsiServerInventoryView `json:"inventory,omitempty"`
}

// DeleteIscsiServerEventView DeleteIscsiServerEvent
type DeleteIscsiServerEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachIscsiServerFromClusterEventView DetachIscsiServerFromClusterEvent
type DetachIscsiServerFromClusterEventView struct {
	Inventory IscsiServerInventoryView `json:"inventory,omitempty"`
}

