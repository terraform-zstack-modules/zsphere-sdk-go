// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateImageParamDetail UpdateImage detail param
type UpdateImageParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	GuestOsType *string `json:"guestOsType,omitempty"`
	MediaType *string `json:"mediaType,omitempty"`
	Format *string `json:"format,omitempty"`
	System *bool `json:"system,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	Virtio *bool `json:"virtio,omitempty"`
}

// UpdateImageParam UpdateImage request param
type UpdateImageParam struct {
	BaseParam
	Params UpdateImageParamDetail `json:"updateImage"`
}
// AddImageParamDetail AddImage detail param
type AddImageParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Url string `json:"url" validate:"required"`
	MediaType *string `json:"mediaType,omitempty"`
	GuestOsType *string `json:"guestOsType,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	System bool `json:"system,omitempty"`
	Format *string `json:"format,omitempty"`
	Platform *string `json:"platform,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids" validate:"required"`
	Type *string `json:"type,omitempty"`
	Virtio *bool `json:"virtio,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddImageParam AddImage request param
type AddImageParam struct {
	BaseParam
	Params AddImageParamDetail `json:"params"`
}
// SyncImageParamDetail SyncImage detail param
type SyncImageParamDetail struct {
}

// SyncImageParam SyncImage request param
type SyncImageParam struct {
	BaseParam
	Params SyncImageParamDetail `json:"syncImage"`
}
// ExpungeImageParamDetail ExpungeImage detail param
type ExpungeImageParamDetail struct {
	Uuid string `json:"uuid,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
}

// ExpungeImageParam ExpungeImage request param
type ExpungeImageParam struct {
	BaseParam
	Params ExpungeImageParamDetail `json:"expungeImage"`
}
// RecoverImageParamDetail RecoverImage detail param
type RecoverImageParamDetail struct {
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
}

// RecoverImageParam RecoverImage request param
type RecoverImageParam struct {
	BaseParam
	Params RecoverImageParamDetail `json:"recoverImage"`
}
// DeleteImageParamDetail DeleteImage detail param
type DeleteImageParamDetail struct {
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteImageParam DeleteImage request param
type DeleteImageParam struct {
	BaseParam
	Params DeleteImageParamDetail `json:"deleteImage"`
}
