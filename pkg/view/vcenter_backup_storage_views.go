// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VCenterBackupStorageInventoryView VCenterBackupStorage
type VCenterBackupStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	VCenterUuid string `json:"vCenterUuid,omitempty"`
	Datastore string `json:"datastore,omitempty"`
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	AttachedZoneUuids []string `json:"attachedZoneUuids,omitempty"`
}

// QueryVCenterBackupStorageView QueryVCenterBackupStorage
type QueryVCenterBackupStorageView struct {
	Inventories []VCenterBackupStorageInventoryView `json:"inventories,omitempty"`
}

