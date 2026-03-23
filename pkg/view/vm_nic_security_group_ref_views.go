// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmNicSecurityGroupRefInventoryView VmNicSecurityGroupRef
type VmNicSecurityGroupRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Priority int `json:"priority,omitempty"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	SecurityGroupUuid string `json:"securityGroupUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
}

// SetVmNicSecurityGroupEventView SetVmNicSecurityGroupEvent
type SetVmNicSecurityGroupEventView struct {
	Inventory []VmNicSecurityGroupRefInventoryView `json:"inventory,omitempty"`
}

// QueryVmNicInSecurityGroupView QueryVmNicInSecurityGroup
type QueryVmNicInSecurityGroupView struct {
	Inventories []VmNicSecurityGroupRefInventoryView `json:"inventories,omitempty"`
}

