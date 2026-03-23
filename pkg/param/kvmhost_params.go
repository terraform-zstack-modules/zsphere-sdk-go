// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddKVMHostParamDetail AddKVMHost detail param
type AddKVMHostParamDetail struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	SshPort *int `json:"sshPort,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddKVMHostParam AddKVMHost request param
type AddKVMHostParam struct {
	BaseParam
	Params AddKVMHostParamDetail `json:"params"`
}
// UpdateKVMHostParamDetail UpdateKVMHost detail param
type UpdateKVMHostParamDetail struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	SshPort *int `json:"sshPort,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ManagementIp *string `json:"managementIp,omitempty"`
}

// UpdateKVMHostParam UpdateKVMHost request param
type UpdateKVMHostParam struct {
	BaseParam
	Params UpdateKVMHostParamDetail `json:"updateKVMHost"`
}
