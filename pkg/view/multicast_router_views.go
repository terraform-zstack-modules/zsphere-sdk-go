// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MulticastRouterInventoryView MulticastRouter
type MulticastRouterInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	RpGroups []MulticastRouterRendezvousPointInventoryView `json:"rpGroups,omitempty"`
	VpcVrs []MulticastRouterVpcVRouterRefInventoryView `json:"vpcVrs,omitempty"`
}

// RemoveRendezvousPointFromMulticastRouterEventView RemoveRendezvousPointFromMulticastRouterEvent
type RemoveRendezvousPointFromMulticastRouterEventView struct {
	Inventory MulticastRouterInventoryView `json:"inventory,omitempty"`
}

// CreateMulticastRouterEventView CreateMulticastRouterEvent
type CreateMulticastRouterEventView struct {
	Inventory MulticastRouterInventoryView `json:"inventory,omitempty"`
}

// QueryMulticastRouterView QueryMulticastRouter
type QueryMulticastRouterView struct {
	Inventories []MulticastRouterInventoryView `json:"inventories,omitempty"`
}

// ChangeMulticastRouterStateEventView ChangeMulticastRouterStateEvent
type ChangeMulticastRouterStateEventView struct {
	Inventory MulticastRouterInventoryView `json:"inventory,omitempty"`
}

// DeleteMulticastRouterEventView DeleteMulticastRouterEvent
type DeleteMulticastRouterEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddRendezvousPointToMulticastRouterEventView AddRendezvousPointToMulticastRouterEvent
type AddRendezvousPointToMulticastRouterEventView struct {
	Inventory MulticastRouterInventoryView `json:"inventory,omitempty"`
}

