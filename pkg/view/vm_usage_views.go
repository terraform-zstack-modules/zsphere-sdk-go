// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmUsageInventoryView VmUsage
type VmUsageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	DateInLong int64 `json:"dateInLong,omitempty"`
	VmUuid string `json:"vmUuid,omitempty"`
	State string `json:"state,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	RootVolumeSize int64 `json:"rootVolumeSize,omitempty"`
	Inventory string `json:"inventory,omitempty"`
}

