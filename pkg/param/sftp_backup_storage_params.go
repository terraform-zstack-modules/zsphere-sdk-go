// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateSftpBackupStorageParamDetail UpdateSftpBackupStorage detail param
type UpdateSftpBackupStorageParamDetail struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	SshPort *int `json:"sshPort,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateSftpBackupStorageParam UpdateSftpBackupStorage request param
type UpdateSftpBackupStorageParam struct {
	BaseParam
	Params UpdateSftpBackupStorageParamDetail `json:"updateSftpBackupStorage"`
}
// ReconnectSftpBackupStorageParamDetail ReconnectSftpBackupStorage detail param
type ReconnectSftpBackupStorageParamDetail struct {
}

// ReconnectSftpBackupStorageParam ReconnectSftpBackupStorage request param
type ReconnectSftpBackupStorageParam struct {
	BaseParam
	Params ReconnectSftpBackupStorageParamDetail `json:"reconnectSftpBackupStorage"`
}
// AddSftpBackupStorageParamDetail AddSftpBackupStorage detail param
type AddSftpBackupStorageParamDetail struct {
	Hostname string `json:"hostname" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	SshPort *int `json:"sshPort,omitempty"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ImportImages *bool `json:"importImages,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSftpBackupStorageParam AddSftpBackupStorage request param
type AddSftpBackupStorageParam struct {
	BaseParam
	Params AddSftpBackupStorageParamDetail `json:"params"`
}
