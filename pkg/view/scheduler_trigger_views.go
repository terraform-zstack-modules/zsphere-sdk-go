// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SchedulerTriggerInventoryView SchedulerTrigger
type SchedulerTriggerInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Cron string `json:"cron,omitempty"`
	SchedulerType string `json:"schedulerType,omitempty"`
	SchedulerInterval int `json:"schedulerInterval,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	StopTime time.Time `json:"stopTime,omitempty"`
	JobsUuid []string `json:"jobsUuid,omitempty"`
	JobGroupsUuid []string `json:"jobGroupsUuid,omitempty"`
}

// GetAvailableTriggersView GetAvailableTriggers
type GetAvailableTriggersView struct {
	Inventories []SchedulerTriggerInventoryView `json:"inventories,omitempty"`
}

// DeleteSchedulerTriggerEventView DeleteSchedulerTriggerEvent
type DeleteSchedulerTriggerEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateSchedulerTriggerEventView CreateSchedulerTriggerEvent
type CreateSchedulerTriggerEventView struct {
	Inventory SchedulerTriggerInventoryView `json:"inventory,omitempty"`
}

// QuerySchedulerTriggerView QuerySchedulerTrigger
type QuerySchedulerTriggerView struct {
	Inventories []SchedulerTriggerInventoryView `json:"inventories,omitempty"`
}

// UpdateSchedulerTriggerEventView UpdateSchedulerTriggerEvent
type UpdateSchedulerTriggerEventView struct {
	Inventory SchedulerTriggerInventoryView `json:"inventory,omitempty"`
}

