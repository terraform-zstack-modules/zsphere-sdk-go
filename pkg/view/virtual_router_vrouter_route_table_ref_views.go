// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VirtualRouterVRouterRouteTableRefInventoryView VirtualRouterVRouterRouteTableRef
type VirtualRouterVRouterRouteTableRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid,omitempty"`
	RouteTableUuid string `json:"routeTableUuid,omitempty"`
}

// QueryVirtualRouterVRouterRouteTableRefView QueryVirtualRouterVRouterRouteTableRef
type QueryVirtualRouterVRouterRouteTableRefView struct {
	Inventories []VirtualRouterVRouterRouteTableRefInventoryView `json:"inventories,omitempty"`
}

