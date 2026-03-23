// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ImageReplicationGroupBackupStorageRefInventoryView ImageReplicationGroupBackupStorageRef
type ImageReplicationGroupBackupStorageRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ReplicationGroupUuid string `json:"replicationGroupUuid,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
}

// AddBackupStoragesToReplicationGroupEventView AddBackupStoragesToReplicationGroupEvent
type AddBackupStoragesToReplicationGroupEventView struct {
	Inventories []ImageReplicationGroupBackupStorageRefInventoryView `json:"inventories,omitempty"`
}

