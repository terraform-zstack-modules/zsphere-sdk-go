// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ImagePackageInventoryView ImagePackage
type ImagePackageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	VmUuid string `json:"vmUuid,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	State string `json:"state,omitempty"`
	ExportUrl string `json:"exportUrl,omitempty"`
	Md5Sum string `json:"md5Sum,omitempty"`
	Format string `json:"format,omitempty"`
	Size int64 `json:"size,omitempty"`
}

// UpdateImagePackageEventView UpdateImagePackageEvent
type UpdateImagePackageEventView struct {
	Inventory ImagePackageInventoryView `json:"inventory,omitempty"`
}

// ExportVmOvaPackageEventView ExportVmOvaPackageEvent
type ExportVmOvaPackageEventView struct {
	Inventory ImagePackageInventoryView `json:"inventory,omitempty"`
}

// QueryImagePackageView QueryImagePackage
type QueryImagePackageView struct {
	Inventories []ImagePackageInventoryView `json:"inventories,omitempty"`
}

// DeleteImagePackageEventView DeleteImagePackageEvent
type DeleteImagePackageEventView struct {
	Success bool `json:"success,omitempty"`
}

