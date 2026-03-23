// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PolicyRouteRuleSetL3RefInventoryView PolicyRouteRuleSetL3Ref
type PolicyRouteRuleSetL3RefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	RuleSetUuid string `json:"ruleSetUuid,omitempty"`
}

// QueryPolicyRouteRuleSetL3RefView QueryPolicyRouteRuleSetL3Ref
type QueryPolicyRouteRuleSetL3RefView struct {
	Inventories []PolicyRouteRuleSetL3RefInventoryView `json:"inventories,omitempty"`
}

