// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SchedulerJobHistoryInventoryView SchedulerJobHistory
type SchedulerJobHistoryInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	TriggerUuid string `json:"triggerUuid,omitempty"`
	SchedulerJobUuid string `json:"schedulerJobUuid,omitempty"`
	SchedulerJobGroupUuid string `json:"schedulerJobGroupUuid,omitempty"`
	JobType string `json:"jobType,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	ExecuteTime int64 `json:"executeTime,omitempty"`
	TargetResourceUuid string `json:"targetResourceUuid,omitempty"`
	RequestDump string `json:"requestDump,omitempty"`
	ResultDump string `json:"resultDump,omitempty"`
	Success bool `json:"success,omitempty"`
	FireInstanceId string `json:"fireInstanceId,omitempty"`
}

// QuerySchedulerJobHistoryView QuerySchedulerJobHistory
type QuerySchedulerJobHistoryView struct {
	Inventories []SchedulerJobHistoryInventoryView `json:"inventories,omitempty"`
}

