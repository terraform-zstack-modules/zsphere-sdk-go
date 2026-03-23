// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AutoScalingVmTemplateInventoryView AutoScalingVmTemplate
type AutoScalingVmTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	VmInstanceName string `json:"vmInstanceName,omitempty"`
	VmInstanceType string `json:"vmInstanceType,omitempty"`
	VmInstanceDescription string `json:"vmInstanceDescription,omitempty"`
	VmInstanceOfferingUuid string `json:"vmInstanceOfferingUuid,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	RootDiskOfferingUuid string `json:"rootDiskOfferingUuid,omitempty"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	VmInstanceZoneUuid string `json:"vmInstanceZoneUuid,omitempty"`
	VmInstanceClusterUuid string `json:"vmInstanceClusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
}

// UpdateAutoScalingTemplateEventView UpdateAutoScalingTemplateEvent
type UpdateAutoScalingTemplateEventView struct {
	Inventory AutoScalingTemplateInventoryView `json:"inventory,omitempty"`
}

// CreateAutoScalingTemplateEventView CreateAutoScalingTemplateEvent
type CreateAutoScalingTemplateEventView struct {
	Inventory AutoScalingTemplateInventoryView `json:"inventory,omitempty"`
}

// QueryAutoScalingVmTemplateView QueryAutoScalingVmTemplate
type QueryAutoScalingVmTemplateView struct {
	Inventories []AutoScalingVmTemplateInventoryView `json:"inventories,omitempty"`
}

