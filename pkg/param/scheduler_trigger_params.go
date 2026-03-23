// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteSchedulerTriggerParamDetail DeleteSchedulerTrigger detail param
type DeleteSchedulerTriggerParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteSchedulerTriggerParam DeleteSchedulerTrigger request param
type DeleteSchedulerTriggerParam struct {
	BaseParam
	Params DeleteSchedulerTriggerParamDetail `json:"deleteSchedulerTrigger"`
}
// CreateSchedulerTriggerParamDetail CreateSchedulerTrigger detail param
type CreateSchedulerTriggerParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	SchedulerInterval *int `json:"schedulerInterval,omitempty"`
	RepeatCount *int `json:"repeatCount,omitempty"`
	StartTime *int64 `json:"startTime,omitempty"`
	StopTime *int64 `json:"stopTime,omitempty"`
	SchedulerType string `json:"schedulerType" validate:"required"`
	Cron *string `json:"cron,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSchedulerTriggerParam CreateSchedulerTrigger request param
type CreateSchedulerTriggerParam struct {
	BaseParam
	Params CreateSchedulerTriggerParamDetail `json:"params"`
}
// UpdateSchedulerTriggerParamDetail UpdateSchedulerTrigger detail param
type UpdateSchedulerTriggerParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	SchedulerInterval *int `json:"schedulerInterval,omitempty"`
	RepeatCount *int `json:"repeatCount,omitempty"`
	StartTime *int64 `json:"startTime,omitempty"`
	StopTime *int64 `json:"stopTime,omitempty"`
	Cron *string `json:"cron,omitempty"`
	SchedulerType *string `json:"schedulerType,omitempty"`
}

// UpdateSchedulerTriggerParam UpdateSchedulerTrigger request param
type UpdateSchedulerTriggerParam struct {
	BaseParam
	Params UpdateSchedulerTriggerParamDetail `json:"updateSchedulerTrigger"`
}
