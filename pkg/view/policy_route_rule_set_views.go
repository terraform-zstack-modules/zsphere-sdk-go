// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PolicyRouteRuleSetInventoryView PolicyRouteRuleSet
type PolicyRouteRuleSetInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	Rules []PolicyRouteRuleInventoryView `json:"rules,omitempty"`
	L3Refs []PolicyRouteRuleSetL3RefInventoryView `json:"l3Refs,omitempty"`
}

// GetPolicyRouteRuleSetFromVirtualRouterView GetPolicyRouteRuleSetFromVirtualRouter
type GetPolicyRouteRuleSetFromVirtualRouterView struct {
	Inventories []PolicyRouteRuleSetInventoryView `json:"inventories,omitempty"`
}

// CreatePolicyRouteRuleSetEventView CreatePolicyRouteRuleSetEvent
type CreatePolicyRouteRuleSetEventView struct {
	Inventory PolicyRouteRuleSetInventoryView `json:"inventory,omitempty"`
}

// UpdatePolicyRouteRuleSetEventView UpdatePolicyRouteRuleSetEvent
type UpdatePolicyRouteRuleSetEventView struct {
	Inventory PolicyRouteRuleSetInventoryView `json:"inventory,omitempty"`
}

// DeletePolicyRouteRuleSetEventView DeletePolicyRouteRuleSetEvent
type DeletePolicyRouteRuleSetEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryPolicyRouteRuleSetView QueryPolicyRouteRuleSet
type QueryPolicyRouteRuleSetView struct {
	Inventories []PolicyRouteRuleSetInventoryView `json:"inventories,omitempty"`
}

