// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ExternalBackupInventoryView ExternalBackup
type ExternalBackupInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	TotalSize int64 `json:"totalSize,omitempty"`
	Version string `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
}

// QueryExternalBackupView QueryExternalBackup
type QueryExternalBackupView struct {
	Inventories []ExternalBackupInventoryView `json:"inventories,omitempty"`
}

// DeleteExternalBackupEventView DeleteExternalBackupEvent
type DeleteExternalBackupEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateExternalBackupEventView CreateExternalBackupEvent
type CreateExternalBackupEventView struct {
	Inventory ExternalBackupInventoryView `json:"inventory,omitempty"`
}

