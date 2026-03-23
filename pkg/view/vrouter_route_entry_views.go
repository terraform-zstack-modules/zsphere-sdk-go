// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VRouterRouteEntryInventoryView VRouterRouteEntry
type VRouterRouteEntryInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	RouteTableUuid string `json:"routeTableUuid,omitempty"`
	Destination string `json:"destination,omitempty"`
	Target string `json:"target,omitempty"`
	Distance int `json:"distance,omitempty"`
}

// QueryVRouterRouteEntryView QueryVRouterRouteEntry
type QueryVRouterRouteEntryView struct {
	Inventories []VRouterRouteEntryInventoryView `json:"inventories,omitempty"`
}

// AddVRouterRouteEntryEventView AddVRouterRouteEntryEvent
type AddVRouterRouteEntryEventView struct {
	Inventory VRouterRouteEntryInventoryView `json:"inventory,omitempty"`
}

