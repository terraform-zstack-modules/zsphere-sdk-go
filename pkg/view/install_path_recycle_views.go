// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// InstallPathRecycleInventoryView InstallPathRecycle
type InstallPathRecycleInventoryView struct {
	BaseInfoView
	BaseTimeView
	TrashId int64 `json:"trashId,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	StorageUuid string `json:"storageUuid,omitempty"`
	StorageType string `json:"storageType,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	IsFolder bool `json:"isFolder,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	Size int64 `json:"size,omitempty"`
	TrashType string `json:"trashType,omitempty"`
}

// GetTrashOnBackupStorageView GetTrashOnBackupStorage
type GetTrashOnBackupStorageView struct {
	Inventories []InstallPathRecycleInventoryView `json:"inventories,omitempty"`
}

// GetTrashOnPrimaryStorageView GetTrashOnPrimaryStorage
type GetTrashOnPrimaryStorageView struct {
	Inventories []InstallPathRecycleInventoryView `json:"inventories,omitempty"`
}

