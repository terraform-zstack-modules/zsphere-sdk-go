// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CpuMemoryCapacityDataView CpuMemoryCapacityData
type CpuMemoryCapacityDataView struct {
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TotalCpu int64 `json:"totalCpu,omitempty"`
	AvailableCpu int64 `json:"availableCpu,omitempty"`
	TotalMemory int64 `json:"totalMemory,omitempty"`
	AvailableMemory int64 `json:"availableMemory,omitempty"`
	ManagedCpuNum int64 `json:"managedCpuNum,omitempty"`
}

