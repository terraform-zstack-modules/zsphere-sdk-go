// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeBackupStorageRefInventoryView VolumeBackupStorageRef
type VolumeBackupStorageRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	VolumeBackupUuid string `json:"volumeBackupUuid,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	Status string `json:"status,omitempty"`
}

