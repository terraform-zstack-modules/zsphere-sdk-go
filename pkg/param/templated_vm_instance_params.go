// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateTemplatedVmInstanceParamDetail UpdateTemplatedVmInstance detail param
type UpdateTemplatedVmInstanceParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
}

// UpdateTemplatedVmInstanceParam UpdateTemplatedVmInstance request param
type UpdateTemplatedVmInstanceParam struct {
	BaseParam
	Params UpdateTemplatedVmInstanceParamDetail `json:"updateTemplatedVmInstance"`
}
// DeleteTemplatedVmInstanceParamDetail DeleteTemplatedVmInstance detail param
type DeleteTemplatedVmInstanceParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteTemplatedVmInstanceParam DeleteTemplatedVmInstance request param
type DeleteTemplatedVmInstanceParam struct {
	BaseParam
	Params DeleteTemplatedVmInstanceParamDetail `json:"deleteTemplatedVmInstance"`
}
