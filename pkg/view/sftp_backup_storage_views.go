// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SftpBackupStorageInventoryView SftpBackupStorage
type SftpBackupStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Hostname string `json:"hostname,omitempty"`
	Username string `json:"username,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	AttachedZoneUuids []string `json:"attachedZoneUuids,omitempty"`
}

// QuerySftpBackupStorageView QuerySftpBackupStorage
type QuerySftpBackupStorageView struct {
	Inventories []SftpBackupStorageInventoryView `json:"inventories,omitempty"`
}

// ReconnectSftpBackupStorageEventView ReconnectSftpBackupStorageEvent
type ReconnectSftpBackupStorageEventView struct {
	Inventory SftpBackupStorageInventoryView `json:"inventory,omitempty"`
}

// AddSftpBackupStorageEventView AddSftpBackupStorageEvent
type AddSftpBackupStorageEventView struct {
	Inventory BackupStorageInventoryView `json:"inventory,omitempty"`
}

