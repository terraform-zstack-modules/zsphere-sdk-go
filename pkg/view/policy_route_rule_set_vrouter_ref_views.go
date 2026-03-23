// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PolicyRouteRuleSetVRouterRefInventoryView PolicyRouteRuleSetVRouterRef
type PolicyRouteRuleSetVRouterRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	VRouterUuid string `json:"vRouterUuid,omitempty"`
	RuleSetUuid string `json:"ruleSetUuid,omitempty"`
}

// QueryPolicyRouteRuleSetVRouterRefView QueryPolicyRouteRuleSetVRouterRef
type QueryPolicyRouteRuleSetVRouterRefView struct {
	Inventories []PolicyRouteRuleSetVRouterRefInventoryView `json:"inventories,omitempty"`
}

