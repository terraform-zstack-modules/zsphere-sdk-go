// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SchedulerJobGroupSchedulerTriggerRefInventoryView SchedulerJobGroupSchedulerTriggerRef
type SchedulerJobGroupSchedulerTriggerRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	SchedulerJobGroupUuid string `json:"schedulerJobGroupUuid,omitempty"`
	SchedulerTriggerUuid string `json:"schedulerTriggerUuid,omitempty"`
}

// AddSchedulerJobGroupToSchedulerTriggerEventView AddSchedulerJobGroupToSchedulerTriggerEvent
type AddSchedulerJobGroupToSchedulerTriggerEventView struct {
	Inventory SchedulerJobGroupSchedulerTriggerRefInventoryView `json:"inventory,omitempty"`
}

