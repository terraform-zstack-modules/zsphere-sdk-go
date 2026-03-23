// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ImageReplicationGroupInventoryView ImageReplicationGroup
type ImageReplicationGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	BackupStorageRefs []ImageReplicationGroupBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
}

// DeleteImageReplicationGroupEventView DeleteImageReplicationGroupEvent
type DeleteImageReplicationGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryImageReplicationGroupView QueryImageReplicationGroup
type QueryImageReplicationGroupView struct {
	Inventories []ImageReplicationGroupInventoryView `json:"inventories,omitempty"`
}

// CreateImageReplicationGroupEventView CreateImageReplicationGroupEvent
type CreateImageReplicationGroupEventView struct {
	Inventory ImageReplicationGroupInventoryView `json:"inventory,omitempty"`
}

