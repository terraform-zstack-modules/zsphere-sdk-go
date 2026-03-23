// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VirtualRouterLoadBalancerRefInventoryView VirtualRouterLoadBalancerRef
type VirtualRouterLoadBalancerRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid,omitempty"`
	LoadBalancerUuid string `json:"loadBalancerUuid,omitempty"`
}

