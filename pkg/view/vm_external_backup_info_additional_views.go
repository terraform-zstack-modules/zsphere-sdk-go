// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmExternalBackupInfoView VmExternalBackupInfo
type VmExternalBackupInfoView struct {
	LiveBackup bool `json:"liveBackup,omitempty"`
	Volumes []VolumeExternalBackupInfoView `json:"volumes,omitempty"`
	TotalSize int64 `json:"totalSize,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
}

