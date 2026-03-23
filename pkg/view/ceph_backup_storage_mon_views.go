// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CephBackupStorageMonInventoryView CephBackupStorageMon
type CephBackupStorageMonInventoryView struct {
	BaseInfoView
	BaseTimeView
	Hostname string `json:"hostname,omitempty"`
	MonPort int `json:"monPort,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	MonAddr string `json:"monAddr,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	Status string `json:"status,omitempty"`
	SshUsername string `json:"sshUsername,omitempty"`
	SshPassword string `json:"sshPassword,omitempty"`
	MonUuid string `json:"monUuid,omitempty"`
}

// UpdateCephBackupStorageMonEventView UpdateCephBackupStorageMonEvent
type UpdateCephBackupStorageMonEventView struct {
	Inventory CephBackupStorageInventoryView `json:"inventory,omitempty"`
}

