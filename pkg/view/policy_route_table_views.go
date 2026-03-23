// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PolicyRouteTableInventoryView PolicyRouteTable
type PolicyRouteTableInventoryView struct {
	BaseInfoView
	BaseTimeView
	TableNumber int `json:"tableNumber,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	Routes []PolicyRouteTableRouteEntryInventoryView `json:"routes,omitempty"`
}

// CreatePolicyRouteTableEventView CreatePolicyRouteTableEvent
type CreatePolicyRouteTableEventView struct {
	Inventory PolicyRouteTableInventoryView `json:"inventory,omitempty"`
}

// DeletePolicyRouteTableEventView DeletePolicyRouteTableEvent
type DeletePolicyRouteTableEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryPolicyRouteTableView QueryPolicyRouteTable
type QueryPolicyRouteTableView struct {
	Inventories []PolicyRouteTableInventoryView `json:"inventories,omitempty"`
}

