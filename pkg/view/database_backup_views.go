// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// DatabaseBackupInventoryView DatabaseBackup
type DatabaseBackupInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Size int64 `json:"size,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	BackupStorageRefs []DatabaseBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
}

// DeleteDatabaseBackupEventView DeleteDatabaseBackupEvent
type DeleteDatabaseBackupEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncDatabaseBackupFromImageStoreBackupStorageEventView SyncDatabaseBackupFromImageStoreBackupStorageEvent
type SyncDatabaseBackupFromImageStoreBackupStorageEventView struct {
	Inventory DatabaseBackupInventoryView `json:"inventory,omitempty"`
}

// CreateDatabaseBackupEventView CreateDatabaseBackupEvent
type CreateDatabaseBackupEventView struct {
	Inventory DatabaseBackupInventoryView `json:"inventory,omitempty"`
}

// QueryDatabaseBackupView QueryDatabaseBackup
type QueryDatabaseBackupView struct {
	Inventories []DatabaseBackupInventoryView `json:"inventories,omitempty"`
}

// SyncDatabaseBackupEventView SyncDatabaseBackupEvent
type SyncDatabaseBackupEventView struct {
	Result SyncBackupResultView `json:"result,omitempty"`
}

