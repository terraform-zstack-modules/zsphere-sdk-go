// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PortMirrorInventoryView PortMirror
type PortMirrorInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	MirrorNetworkUuid string `json:"mirrorNetworkUuid,omitempty"`
	Sessions []PortMirrorSessionInventoryView `json:"sessions,omitempty"`
}

// ChangePortMirrorStateEventView ChangePortMirrorStateEvent
type ChangePortMirrorStateEventView struct {
	Inventory PortMirrorInventoryView `json:"inventory,omitempty"`
}

// QueryPortMirrorView QueryPortMirror
type QueryPortMirrorView struct {
	Inventories []PortMirrorInventoryView `json:"inventories,omitempty"`
}

// CreatePortMirrorEventView CreatePortMirrorEvent
type CreatePortMirrorEventView struct {
	Inventory PortMirrorInventoryView `json:"inventory,omitempty"`
}

// DeletePortMirrorEventView DeletePortMirrorEvent
type DeletePortMirrorEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdatePortMirrorEventView UpdatePortMirrorEvent
type UpdatePortMirrorEventView struct {
	Inventory PortMirrorInventoryView `json:"inventory,omitempty"`
}

