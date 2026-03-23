// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PortMirrorSessionInventoryView PortMirrorSession
type PortMirrorSessionInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Status string `json:"status,omitempty"`
	InternalId int64 `json:"internalId,omitempty"`
	SrcEndPoint string `json:"srcEndPoint,omitempty"`
	Type string `json:"type,omitempty"`
	DstEndPoint string `json:"dstEndPoint,omitempty"`
	PortMirrorUuid string `json:"portMirrorUuid,omitempty"`
}

// CreatePortMirrorSessionEventView CreatePortMirrorSessionEvent
type CreatePortMirrorSessionEventView struct {
	Inventory PortMirrorSessionInventoryView `json:"inventory,omitempty"`
}

// DeletePortMirrorSessionEventView DeletePortMirrorSessionEvent
type DeletePortMirrorSessionEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryPortMirrorSessionView QueryPortMirrorSession
type QueryPortMirrorSessionView struct {
	Inventories []PortMirrorSessionInventoryView `json:"inventories,omitempty"`
}

