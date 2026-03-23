// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SlbLoadBalancerInventoryView SlbLoadBalancer
type SlbLoadBalancerInventoryView struct {
	BaseInfoView
	BaseTimeView
	SlbGroupUuid string `json:"slbGroupUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ServerGroupUuid string `json:"serverGroupUuid,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	Listeners []LoadBalancerListenerInventoryView `json:"listeners,omitempty"`
}

