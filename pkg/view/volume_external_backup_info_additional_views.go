// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeExternalBackupInfoView VolumeExternalBackupInfo
type VolumeExternalBackupInfoView struct {
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Size int64 `json:"size,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
}

