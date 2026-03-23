// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddXDragonHostParamDetail AddXDragonHost detail param
type AddXDragonHostParamDetail struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	CpuNum *int `json:"cpuNum,omitempty"`
	CpuSockets *int `json:"cpuSockets,omitempty"`
	TotalPhysicalMemory *int64 `json:"totalPhysicalMemory,omitempty"`
	SshPort *int `json:"sshPort,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddXDragonHostParam AddXDragonHost request param
type AddXDragonHostParam struct {
	BaseParam
	Params AddXDragonHostParamDetail `json:"params"`
}
