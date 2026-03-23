// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateCephBackupStorageMonParamDetail UpdateCephBackupStorageMon detail param
type UpdateCephBackupStorageMonParamDetail struct {
	Hostname *string `json:"hostname,omitempty"`
	SshUsername *string `json:"sshUsername,omitempty"`
	SshPassword *string `json:"sshPassword,omitempty"`
	SshPort *int `json:"sshPort,omitempty"`
	MonPort *int `json:"monPort,omitempty"`
}

// UpdateCephBackupStorageMonParam UpdateCephBackupStorageMon request param
type UpdateCephBackupStorageMonParam struct {
	BaseParam
	Params UpdateCephBackupStorageMonParamDetail `json:"updateCephBackupStorageMon"`
}
