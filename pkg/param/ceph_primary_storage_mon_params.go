// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateCephPrimaryStorageMonParamDetail UpdateCephPrimaryStorageMon detail param
type UpdateCephPrimaryStorageMonParamDetail struct {
	Hostname *string `json:"hostname,omitempty"`
	SshUsername *string `json:"sshUsername,omitempty"`
	SshPassword *string `json:"sshPassword,omitempty"`
	SshPort *int `json:"sshPort,omitempty"`
	MonPort *int `json:"monPort,omitempty"`
}

// UpdateCephPrimaryStorageMonParam UpdateCephPrimaryStorageMon request param
type UpdateCephPrimaryStorageMonParam struct {
	BaseParam
	Params UpdateCephPrimaryStorageMonParamDetail `json:"updateCephPrimaryStorageMon"`
}
