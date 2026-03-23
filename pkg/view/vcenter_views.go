// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VCenterInventoryView VCenter
type VCenterInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	DomainName string `json:"domainName,omitempty"`
	Port int `json:"port,omitempty"`
	UserName string `json:"userName,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Version string `json:"version,omitempty"`
	Https bool `json:"https,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
}

// AddVCenterEventView AddVCenterEvent
type AddVCenterEventView struct {
	Inventory VCenterInventoryView `json:"inventory,omitempty"`
	SkippedResources SkippedResourcesView `json:"skippedResources,omitempty"`
}

// SyncVCenterEventView SyncVCenterEvent
type SyncVCenterEventView struct {
	SkippedResources SkippedResourcesView `json:"skippedResources,omitempty"`
}

// QueryVCenterView QueryVCenter
type QueryVCenterView struct {
	Inventories []VCenterInventoryView `json:"inventories,omitempty"`
}

// UpdateVCenterEventView UpdateVCenterEvent
type UpdateVCenterEventView struct {
	Inventory VCenterInventoryView `json:"inventory,omitempty"`
}

// DeleteVCenterEventView DeleteVCenterEvent
type DeleteVCenterEventView struct {
	Success bool `json:"success,omitempty"`
}

