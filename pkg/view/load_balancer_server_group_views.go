// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LoadBalancerServerGroupInventoryView LoadBalancerServerGroup
type LoadBalancerServerGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	LoadBalancerUuid string `json:"loadBalancerUuid,omitempty"`
	ListenerServerGroupRefs []LoadBalancerListenerServerGroupRefInventoryView `json:"listenerServerGroupRefs,omitempty"`
	ServerIps []LoadBalancerServerGroupServerIpInventoryView `json:"serverIps,omitempty"`
	VmNicRefs []LoadBalancerServerGroupVmNicRefInventoryView `json:"vmNicRefs,omitempty"`
}

// DeleteLoadBalancerServerGroupEventView DeleteLoadBalancerServerGroupEvent
type DeleteLoadBalancerServerGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateLoadBalancerServerGroupEventView CreateLoadBalancerServerGroupEvent
type CreateLoadBalancerServerGroupEventView struct {
	Inventory LoadBalancerServerGroupInventoryView `json:"inventory,omitempty"`
}

// RemoveBackendServerFromServerGroupEventView RemoveBackendServerFromServerGroupEvent
type RemoveBackendServerFromServerGroupEventView struct {
	Inventory LoadBalancerServerGroupInventoryView `json:"inventory,omitempty"`
}

// UpdateLoadBalancerServerGroupEventView UpdateLoadBalancerServerGroupEvent
type UpdateLoadBalancerServerGroupEventView struct {
	Inventory LoadBalancerServerGroupInventoryView `json:"inventory,omitempty"`
}

// ChangeLoadBalancerBackendServerEventView ChangeLoadBalancerBackendServerEvent
type ChangeLoadBalancerBackendServerEventView struct {
	Inventory LoadBalancerServerGroupInventoryView `json:"inventory,omitempty"`
}

// QueryLoadBalancerServerGroupView QueryLoadBalancerServerGroup
type QueryLoadBalancerServerGroupView struct {
	Inventories []LoadBalancerServerGroupInventoryView `json:"inventories,omitempty"`
}

// AddBackendServerToServerGroupEventView AddBackendServerToServerGroupEvent
type AddBackendServerToServerGroupEventView struct {
	Inventory LoadBalancerServerGroupInventoryView `json:"inventory,omitempty"`
}

