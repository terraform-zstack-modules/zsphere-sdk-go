// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddSimulatorHostParamDetail AddSimulatorHost detail param
type AddSimulatorHostParamDetail struct {
	MemoryCapacity int64 `json:"memoryCapacity" validate:"required"`
	CpuCapacity int64 `json:"cpuCapacity" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSimulatorHostParam AddSimulatorHost request param
type AddSimulatorHostParam struct {
	BaseParam
	Params AddSimulatorHostParamDetail `json:"params"`
}
