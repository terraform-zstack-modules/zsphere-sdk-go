// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VRouterRouteTableInventoryView VRouterRouteTable
type VRouterRouteTableInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	AttachedRouterRefs []VirtualRouterVRouterRouteTableRefInventoryView `json:"attachedRouterRefs,omitempty"`
	RouteEntries []VRouterRouteEntryInventoryView `json:"routeEntries,omitempty"`
}

// QueryVRouterRouteTableView QueryVRouterRouteTable
type QueryVRouterRouteTableView struct {
	Inventories []VRouterRouteTableInventoryView `json:"inventories,omitempty"`
}

// GetVRouterRouteTableView GetVRouterRouteTable
type GetVRouterRouteTableView struct {
	Inventories []VRouterRouteEntryAOView `json:"inventories,omitempty"`
}

// CreateVRouterRouteTableEventView CreateVRouterRouteTableEvent
type CreateVRouterRouteTableEventView struct {
	Inventory VRouterRouteTableInventoryView `json:"inventory,omitempty"`
}

// DeleteVRouterRouteEntryEventView DeleteVRouterRouteEntryEvent
type DeleteVRouterRouteEntryEventView struct {
	Inventory VRouterRouteTableInventoryView `json:"inventory,omitempty"`
}

// UpdateVRouterRouteTableEventView UpdateVRouterRouteTableEvent
type UpdateVRouterRouteTableEventView struct {
	Inventory VRouterRouteTableInventoryView `json:"inventory,omitempty"`
}

// DetachVRouterRouteTableFromVRouterEventView DetachVRouterRouteTableFromVRouterEvent
type DetachVRouterRouteTableFromVRouterEventView struct {
	Inventory VRouterRouteTableInventoryView `json:"inventory,omitempty"`
}

// DeleteVRouterRouteTableEventView DeleteVRouterRouteTableEvent
type DeleteVRouterRouteTableEventView struct {
	Success bool `json:"success,omitempty"`
}

// AttachVRouterRouteTableToVRouterEventView AttachVRouterRouteTableToVRouterEvent
type AttachVRouterRouteTableToVRouterEventView struct {
	Inventory VRouterRouteTableInventoryView `json:"inventory,omitempty"`
}

