// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VirtualRouterSoftwareVersionInventoryView VirtualRouterSoftwareVersion
type VirtualRouterSoftwareVersionInventoryView struct {
	BaseInfoView
	BaseTimeView
	SoftwareName string `json:"softwareName,omitempty"`
	CurrentVersion string `json:"currentVersion,omitempty"`
	LatestVersion string `json:"latestVersion,omitempty"`
}

// GetVirtualRouterSoftwareVersionView GetVirtualRouterSoftwareVersion
type GetVirtualRouterSoftwareVersionView struct {
	Inventories []VirtualRouterSoftwareVersionInventoryView `json:"inventories,omitempty"`
}

// UpdateVirtualRouterSoftwareVersionEventView UpdateVirtualRouterSoftwareVersionEvent
type UpdateVirtualRouterSoftwareVersionEventView struct {
	Success bool `json:"success,omitempty"`
}

