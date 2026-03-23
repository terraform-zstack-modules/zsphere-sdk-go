// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SchedulerJobSchedulerTriggerInventoryView SchedulerJobSchedulerTrigger
type SchedulerJobSchedulerTriggerInventoryView struct {
	BaseInfoView
	BaseTimeView
	SchedulerJobUuid string `json:"schedulerJobUuid,omitempty"`
	SchedulerTriggerUuid string `json:"schedulerTriggerUuid,omitempty"`
}

// AddSchedulerJobToSchedulerTriggerEventView AddSchedulerJobToSchedulerTriggerEvent
type AddSchedulerJobToSchedulerTriggerEventView struct {
	Inventory SchedulerJobSchedulerTriggerInventoryView `json:"inventory,omitempty"`
}

