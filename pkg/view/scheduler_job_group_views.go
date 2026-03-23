// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SchedulerJobGroupInventoryView SchedulerJobGroup
type SchedulerJobGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	JobType string `json:"jobType,omitempty"`
	JobData string `json:"jobData,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	TriggersUuid []string `json:"triggersUuid,omitempty"`
	JobsUuid []string `json:"jobsUuid,omitempty"`
}

// CreateSchedulerJobGroupEventView CreateSchedulerJobGroupEvent
type CreateSchedulerJobGroupEventView struct {
	Inventory SchedulerJobGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteSchedulerJobGroupEventView DeleteSchedulerJobGroupEvent
type DeleteSchedulerJobGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateSchedulerJobGroupEventView UpdateSchedulerJobGroupEvent
type UpdateSchedulerJobGroupEventView struct {
	Inventory SchedulerJobGroupInventoryView `json:"inventory,omitempty"`
}

// QuerySchedulerJobGroupView QuerySchedulerJobGroup
type QuerySchedulerJobGroupView struct {
	Inventories []SchedulerJobGroupInventoryView `json:"inventories,omitempty"`
}

