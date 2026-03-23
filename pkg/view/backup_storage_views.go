// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BackupStorageInventoryView BackupStorage
type BackupStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	AttachedZoneUuids []string `json:"attachedZoneUuids,omitempty"`
}

// ChangeBackupStorageStateEventView ChangeBackupStorageStateEvent
type ChangeBackupStorageStateEventView struct {
	Inventory BackupStorageInventoryView `json:"inventory,omitempty"`
}

// QueryBackupStorageView QueryBackupStorage
type QueryBackupStorageView struct {
	Inventories []BackupStorageInventoryView `json:"inventories,omitempty"`
}

// DetachBackupStorageFromZoneEventView DetachBackupStorageFromZoneEvent
type DetachBackupStorageFromZoneEventView struct {
	Inventory BackupStorageInventoryView `json:"inventory,omitempty"`
}

// AddBackupStorageEventView AddBackupStorageEvent
type AddBackupStorageEventView struct {
	Inventory BackupStorageInventoryView `json:"inventory,omitempty"`
}

// GetCandidateBackupStorageForCreatingImageView GetCandidateBackupStorageForCreatingImage
type GetCandidateBackupStorageForCreatingImageView struct {
	Inventories []BackupStorageInventoryView `json:"inventories,omitempty"`
}

// UpdateBackupStorageEventView UpdateBackupStorageEvent
type UpdateBackupStorageEventView struct {
	Inventory BackupStorageInventoryView `json:"inventory,omitempty"`
}

// DeleteBackupStorageEventView DeleteBackupStorageEvent
type DeleteBackupStorageEventView struct {
	Success bool `json:"success,omitempty"`
}

// ReconnectBackupStorageEventView ReconnectBackupStorageEvent
type ReconnectBackupStorageEventView struct {
	Inventory BackupStorageInventoryView `json:"inventory,omitempty"`
}

// AttachBackupStorageToZoneEventView AttachBackupStorageToZoneEvent
type AttachBackupStorageToZoneEventView struct {
	Inventory BackupStorageInventoryView `json:"inventory,omitempty"`
}

// GetBackupStorageCandidatesForImageMigrationView GetBackupStorageCandidatesForImageMigration
type GetBackupStorageCandidatesForImageMigrationView struct {
	Inventories []BackupStorageInventoryView `json:"inventories,omitempty"`
}

