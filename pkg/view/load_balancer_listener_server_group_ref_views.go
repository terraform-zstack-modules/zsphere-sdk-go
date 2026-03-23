// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LoadBalancerListenerServerGroupRefInventoryView LoadBalancerListenerServerGroupRef
type LoadBalancerListenerServerGroupRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	ListenerUuid string `json:"listenerUuid,omitempty"`
	ServerGroupUuid string `json:"serverGroupUuid,omitempty"`
}

