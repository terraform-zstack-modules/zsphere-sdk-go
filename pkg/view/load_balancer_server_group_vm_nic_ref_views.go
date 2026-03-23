// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LoadBalancerServerGroupVmNicRefInventoryView LoadBalancerServerGroupVmNicRef
type LoadBalancerServerGroupVmNicRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	ServerGroupUuid string `json:"serverGroupUuid,omitempty"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	Weight int64 `json:"weight,omitempty"`
	Status string `json:"status,omitempty"`
}

