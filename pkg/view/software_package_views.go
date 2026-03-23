// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SoftwarePackageInventoryView SoftwarePackage
type SoftwarePackageInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostUuid string `json:"hostUuid,omitempty"`
	ManagementNodeUuid string `json:"managementNodeUuid,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	UnzipInstallPath string `json:"unzipInstallPath,omitempty"`
	Type string `json:"type,omitempty"`
	Md5sum string `json:"md5sum,omitempty"`
	Status string `json:"status,omitempty"`
	Size int64 `json:"size,omitempty"`
}

// QuerySoftwarePackageView QuerySoftwarePackage
type QuerySoftwarePackageView struct {
	Inventories []SoftwarePackageInventoryView `json:"inventories,omitempty"`
}

// CleanSoftwarePackageEventView CleanSoftwarePackageEvent
type CleanSoftwarePackageEventView struct {
	Success bool `json:"success,omitempty"`
}

