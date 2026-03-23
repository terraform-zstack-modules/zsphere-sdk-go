// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ExternalBackupStorageInventoryView ExternalBackupStorage
type ExternalBackupStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Identity string `json:"identity,omitempty"`
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	AttachedZoneUuids []string `json:"attachedZoneUuids,omitempty"`
}

// AddExternalBackupStorageEventView AddExternalBackupStorageEvent
type AddExternalBackupStorageEventView struct {
	Inventory ExternalBackupStorageInventoryView `json:"inventory,omitempty"`
}

