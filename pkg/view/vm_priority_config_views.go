// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmPriorityConfigInventoryView VmPriorityConfig
type VmPriorityConfigInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccountUuid string `json:"accountUuid,omitempty"`
	Level string `json:"level,omitempty"`
	CpuShares int `json:"cpuShares,omitempty"`
	OomScoreAdj int `json:"oomScoreAdj,omitempty"`
}

// QueryVmPriorityConfigView QueryVmPriorityConfig
type QueryVmPriorityConfigView struct {
	Inventories []VmPriorityConfigInventoryView `json:"inventories,omitempty"`
}

