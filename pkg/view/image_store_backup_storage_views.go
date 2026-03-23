// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ImageStoreBackupStorageInventoryView ImageStoreBackupStorage
type ImageStoreBackupStorageInventoryView struct {
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

// AddImageStoreBackupStorageEventView AddImageStoreBackupStorageEvent
type AddImageStoreBackupStorageEventView struct {
	Inventory ImageStoreBackupStorageInventoryView `json:"inventory,omitempty"`
}

// QueryImageStoreBackupStorageView QueryImageStoreBackupStorage
type QueryImageStoreBackupStorageView struct {
	Inventories []ImageStoreBackupStorageInventoryView `json:"inventories,omitempty"`
}

// ReconnectImageStoreBackupStorageEventView ReconnectImageStoreBackupStorageEvent
type ReconnectImageStoreBackupStorageEventView struct {
	Inventory ImageStoreBackupStorageInventoryView `json:"inventory,omitempty"`
}

