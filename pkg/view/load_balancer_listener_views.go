// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LoadBalancerListenerInventoryView LoadBalancerListener
type LoadBalancerListenerInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	LoadBalancerUuid string `json:"loadBalancerUuid,omitempty"`
	InstancePort int `json:"instancePort,omitempty"`
	LoadBalancerPort int `json:"loadBalancerPort,omitempty"`
	SecurityPolicyType string `json:"securityPolicyType,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	ServerGroupUuid string `json:"serverGroupUuid,omitempty"`
	VmNicRefs []LoadBalancerListenerVmNicRefInventoryView `json:"vmNicRefs,omitempty"`
	AclRefs []LoadBalancerListenerACLRefInventoryView `json:"aclRefs,omitempty"`
	CertificateRefs []LoadBalancerListenerCertificateRefInventoryView `json:"certificateRefs,omitempty"`
	ServerGroupRefs []LoadBalancerListenerServerGroupRefInventoryView `json:"serverGroupRefs,omitempty"`
}

// AddAccessControlListToLoadBalancerEventView AddAccessControlListToLoadBalancerEvent
type AddAccessControlListToLoadBalancerEventView struct {
	Inventory LoadBalancerListenerInventoryView `json:"inventory,omitempty"`
}

// RemoveServerGroupFromLoadBalancerListenerEventView RemoveServerGroupFromLoadBalancerListenerEvent
type RemoveServerGroupFromLoadBalancerListenerEventView struct {
	Inventory LoadBalancerListenerInventoryView `json:"inventory,omitempty"`
}

// RemoveCertificateFromLoadBalancerListenerEventView RemoveCertificateFromLoadBalancerListenerEvent
type RemoveCertificateFromLoadBalancerListenerEventView struct {
	Inventory LoadBalancerListenerInventoryView `json:"inventory,omitempty"`
}

// AddVmNicToLoadBalancerEventView AddVmNicToLoadBalancerEvent
type AddVmNicToLoadBalancerEventView struct {
	Inventory LoadBalancerListenerInventoryView `json:"inventory,omitempty"`
}

// QueryLoadBalancerListenerView QueryLoadBalancerListener
type QueryLoadBalancerListenerView struct {
	Inventories []LoadBalancerListenerInventoryView `json:"inventories,omitempty"`
}

// UpdateLoadBalancerListenerEventView UpdateLoadBalancerListenerEvent
type UpdateLoadBalancerListenerEventView struct {
	Inventory LoadBalancerListenerInventoryView `json:"inventory,omitempty"`
}

// ChangeLoadBalancerListenerEventView ChangeLoadBalancerListenerEvent
type ChangeLoadBalancerListenerEventView struct {
	Inventory LoadBalancerListenerInventoryView `json:"inventory,omitempty"`
}

// CreateLoadBalancerListenerEventView CreateLoadBalancerListenerEvent
type CreateLoadBalancerListenerEventView struct {
	Inventory LoadBalancerListenerInventoryView `json:"inventory,omitempty"`
}

// RemoveAccessControlListFromLoadBalancerEventView RemoveAccessControlListFromLoadBalancerEvent
type RemoveAccessControlListFromLoadBalancerEventView struct {
	Inventory LoadBalancerListenerInventoryView `json:"inventory,omitempty"`
}

// AddCertificateToLoadBalancerListenerEventView AddCertificateToLoadBalancerListenerEvent
type AddCertificateToLoadBalancerListenerEventView struct {
	Inventory LoadBalancerListenerInventoryView `json:"inventory,omitempty"`
}

// AddServerGroupToLoadBalancerListenerEventView AddServerGroupToLoadBalancerListenerEvent
type AddServerGroupToLoadBalancerListenerEventView struct {
	Inventory LoadBalancerListenerInventoryView `json:"inventory,omitempty"`
}

