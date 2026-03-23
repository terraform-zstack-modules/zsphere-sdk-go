// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZBoxBackupInventoryView ZBoxBackup
type ZBoxBackupInventoryView struct {
	BaseInfoView
	BaseTimeView
	ZBoxUuid string `json:"zBoxUuid,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	TotalSize int64 `json:"totalSize,omitempty"`
	Version string `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
}

// QueryZBoxBackupView QueryZBoxBackup
type QueryZBoxBackupView struct {
	Inventories []ZBoxBackupInventoryView `json:"inventories,omitempty"`
}

