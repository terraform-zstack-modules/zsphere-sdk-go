// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ImageInventoryView Image
type ImageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Size int64 `json:"size,omitempty"`
	ActualSize int64 `json:"actualSize,omitempty"`
	Md5Sum string `json:"md5Sum,omitempty"`
	Url string `json:"url,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
	Type string `json:"type,omitempty"`
	Platform string `json:"platform,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Format string `json:"format,omitempty"`
	System bool `json:"system,omitempty"`
	Virtio bool `json:"virtio,omitempty"`
	BackupStorageRefs []ImageBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
	SystemTags []SystemTagInventoryView `json:"systemTags,omitempty"`
}

// GetCandidateIsoForAttachingVmView GetCandidateIsoForAttachingVm
type GetCandidateIsoForAttachingVmView struct {
	Inventories []ImageInventoryView `json:"inventories,omitempty"`
}

// UpdateImageEventView UpdateImageEvent
type UpdateImageEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// GetImageCandidatesForVmToChangeView GetImageCandidatesForVmToChange
type GetImageCandidatesForVmToChangeView struct {
	Inventories []ImageInventoryView `json:"inventories,omitempty"`
}

// ChangeImageStateEventView ChangeImageStateEvent
type ChangeImageStateEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// AddImageEventView AddImageEvent
type AddImageEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// QueryImageView QueryImage
type QueryImageView struct {
	Inventories []ImageInventoryView `json:"inventories,omitempty"`
}

// CreateDataVolumeTemplateFromVolumeEventView CreateDataVolumeTemplateFromVolumeEvent
type CreateDataVolumeTemplateFromVolumeEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// CreateRootVolumeTemplateFromVolumeBackupEventView CreateRootVolumeTemplateFromVolumeBackupEvent
type CreateRootVolumeTemplateFromVolumeBackupEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// CreateRootVolumeTemplateFromVolumeSnapshotEventView CreateRootVolumeTemplateFromVolumeSnapshotEvent
type CreateRootVolumeTemplateFromVolumeSnapshotEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
	Failures []FailureView `json:"failures,omitempty"`
}

// SyncImageEventView SyncImageEvent
type SyncImageEventView struct {
	Success bool `json:"success,omitempty"`
}

// RecoveryImageFromImageStoreBackupStorageEventView RecoveryImageFromImageStoreBackupStorageEvent
type RecoveryImageFromImageStoreBackupStorageEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// GetCandidateImagesForCreatingVmView GetCandidateImagesForCreatingVm
type GetCandidateImagesForCreatingVmView struct {
	Inventories []ImageInventoryView `json:"inventories,omitempty"`
}

// ExpungeImageEventView ExpungeImageEvent
type ExpungeImageEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncImageFromImageStoreBackupStorageEventView SyncImageFromImageStoreBackupStorageEvent
type SyncImageFromImageStoreBackupStorageEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// BackupStorageMigrateImageEventView BackupStorageMigrateImageEvent
type BackupStorageMigrateImageEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// RecoverImageEventView RecoverImageEvent
type RecoverImageEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// CalculateImageHashEventView CalculateImageHashEvent
type CalculateImageHashEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// CreateDataVolumeTemplateFromVolumeSnapshotEventView CreateDataVolumeTemplateFromVolumeSnapshotEvent
type CreateDataVolumeTemplateFromVolumeSnapshotEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
	Failures []FailureView `json:"failures,omitempty"`
}

// DeleteImageEventView DeleteImageEvent
type DeleteImageEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateRootVolumeTemplateFromRootVolumeEventView CreateRootVolumeTemplateFromRootVolumeEvent
type CreateRootVolumeTemplateFromRootVolumeEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// CreateDataVolumeTemplateFromVolumeBackupEventView CreateDataVolumeTemplateFromVolumeBackupEvent
type CreateDataVolumeTemplateFromVolumeBackupEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// SyncImageSizeEventView SyncImageSizeEvent
type SyncImageSizeEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

