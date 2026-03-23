// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateImageStoreBackupStorageParamDetail UpdateImageStoreBackupStorage detail param
type UpdateImageStoreBackupStorageParamDetail struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	SshPort *int `json:"sshPort,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateImageStoreBackupStorageParam UpdateImageStoreBackupStorage request param
type UpdateImageStoreBackupStorageParam struct {
	BaseParam
	Params UpdateImageStoreBackupStorageParamDetail `json:"updateImageStoreBackupStorage"`
}
// ReconnectImageStoreBackupStorageParamDetail ReconnectImageStoreBackupStorage detail param
type ReconnectImageStoreBackupStorageParamDetail struct {
}

// ReconnectImageStoreBackupStorageParam ReconnectImageStoreBackupStorage request param
type ReconnectImageStoreBackupStorageParam struct {
	BaseParam
	Params ReconnectImageStoreBackupStorageParamDetail `json:"reconnectImageStoreBackupStorage"`
}
// AddImageStoreBackupStorageParamDetail AddImageStoreBackupStorage detail param
type AddImageStoreBackupStorageParamDetail struct {
	Hostname string `json:"hostname" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password *string `json:"password,omitempty"`
	SshPort *int `json:"sshPort,omitempty"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ImportImages *bool `json:"importImages,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddImageStoreBackupStorageParam AddImageStoreBackupStorage request param
type AddImageStoreBackupStorageParam struct {
	BaseParam
	Params AddImageStoreBackupStorageParamDetail `json:"params"`
}
