// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SchedulerJobGroupJobRefInventoryView SchedulerJobGroupJobRef
type SchedulerJobGroupJobRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	SchedulerJobGroupUuid string `json:"schedulerJobGroupUuid,omitempty"`
	SchedulerJobUuid string `json:"schedulerJobUuid,omitempty"`
	Priority int `json:"priority,omitempty"`
}

// AddSchedulerJobsToSchedulerJobGroupEventView AddSchedulerJobsToSchedulerJobGroupEvent
type AddSchedulerJobsToSchedulerJobGroupEventView struct {
	Inventories []SchedulerJobGroupJobRefInventoryView `json:"inventories,omitempty"`
}

