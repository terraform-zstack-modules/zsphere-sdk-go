// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmNicSecurityPolicyInventoryView VmNicSecurityPolicy
type VmNicSecurityPolicyInventoryView struct {
	BaseInfoView
	BaseTimeView
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	IngressPolicy string `json:"ingressPolicy,omitempty"`
	EgressPolicy string `json:"egressPolicy,omitempty"`
}

// QueryVmNicSecurityPolicyView QueryVmNicSecurityPolicy
type QueryVmNicSecurityPolicyView struct {
	Inventories []VmNicSecurityPolicyInventoryView `json:"inventories,omitempty"`
}

// ChangeVmNicSecurityPolicyEventView ChangeVmNicSecurityPolicyEvent
type ChangeVmNicSecurityPolicyEventView struct {
	Inventory VmNicSecurityPolicyInventoryView `json:"inventory,omitempty"`
}

