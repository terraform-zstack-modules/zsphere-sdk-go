// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmInstanceResourceMetadataGroupInventoryView VmInstanceResourceMetadataGroup
type VmInstanceResourceMetadataGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	ResourceUuid string `json:"resourceUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	AddressList []VmInstanceResourceMetadataArchiveInventoryView `json:"addressList,omitempty"`
}

// QueryVmInstanceResourceMetadataGroupView QueryVmInstanceResourceMetadataGroup
type QueryVmInstanceResourceMetadataGroupView struct {
	Inventories []VmInstanceResourceMetadataGroupInventoryView `json:"inventories,omitempty"`
}

