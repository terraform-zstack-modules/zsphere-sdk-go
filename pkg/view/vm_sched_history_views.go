// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmSchedHistoryInventoryView VmSchedHistory
type VmSchedHistoryInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	SchedType string `json:"schedType,omitempty"`
	SchedReason string `json:"schedReason,omitempty"`
	FailReason string `json:"failReason,omitempty"`
	Success bool `json:"success,omitempty"`
	LastHostUuid string `json:"lastHostUuid,omitempty"`
	DestHostUuid string `json:"destHostUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
}

// QueryVmSchedHistoryView QueryVmSchedHistory
type QueryVmSchedHistoryView struct {
	Inventories []VmSchedHistoryInventoryView `json:"inventories,omitempty"`
}

