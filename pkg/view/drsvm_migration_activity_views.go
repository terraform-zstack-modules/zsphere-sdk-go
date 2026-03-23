// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// DRSVmMigrationActivityInventoryView DRSVmMigrationActivity
type DRSVmMigrationActivityInventoryView struct {
	BaseInfoView
	BaseTimeView
	DrsUuid string `json:"drsUuid,omitempty"`
	VmUuid string `json:"vmUuid,omitempty"`
	VmSourceHostUuid string `json:"vmSourceHostUuid,omitempty"`
	VmTargetHostUuid string `json:"vmTargetHostUuid,omitempty"`
	Status string `json:"status,omitempty"`
	Result string `json:"result,omitempty"`
	Reason string `json:"reason,omitempty"`
	AdviceUuid string `json:"adviceUuid,omitempty"`
	Cause string `json:"cause,omitempty"`
	EndDate time.Time `json:"endDate,omitempty"`
}

// QueryDRSVmMigrationActivityView QueryDRSVmMigrationActivity
type QueryDRSVmMigrationActivityView struct {
	Inventories []DRSVmMigrationActivityInventoryView `json:"inventories,omitempty"`
}

