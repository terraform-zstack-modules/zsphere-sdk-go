// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmDnsInventoryView VmDns
type VmDnsInventoryView struct {
	BaseInfoView
	BaseTimeView
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	Dns string `json:"dns,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
}

// GetVmDnsView GetVmDns
type GetVmDnsView struct {
	VmDnsList []VmDnsInventoryView `json:"vmDnsList,omitempty"`
	DnsList []string `json:"dnsList,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SetVmDnsEventView SetVmDnsEvent
type SetVmDnsEventView struct {
	Inventories []VmDnsInventoryView `json:"inventories,omitempty"`
}

