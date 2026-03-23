// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AlarmInventoryView Alarm
type AlarmInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	Period int `json:"period,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	MetricName string `json:"metricName,omitempty"`
	Threshold float64 `json:"threshold,omitempty"`
	RepeatInterval int `json:"repeatInterval,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	Status string `json:"status,omitempty"`
	State string `json:"state,omitempty"`
	EnableRecovery bool `json:"enableRecovery,omitempty"`
	Labels []AlarmLabelInventoryView `json:"labels,omitempty"`
	Actions []AlarmActionInventoryView `json:"actions,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
}

// UpdateAlarmEventView UpdateAlarmEvent
type UpdateAlarmEventView struct {
	Inventory AlarmInventoryView `json:"inventory,omitempty"`
}

// RemoveActionFromAlarmEventView RemoveActionFromAlarmEvent
type RemoveActionFromAlarmEventView struct {
	Inventory AlarmInventoryView `json:"inventory,omitempty"`
}

// ChangeAlarmStateEventView ChangeAlarmStateEvent
type ChangeAlarmStateEventView struct {
	Inventory AlarmInventoryView `json:"inventory,omitempty"`
}

// RemoveLabelFromAlarmEventView RemoveLabelFromAlarmEvent
type RemoveLabelFromAlarmEventView struct {
	Inventory AlarmInventoryView `json:"inventory,omitempty"`
}

// DeleteAlarmEventView DeleteAlarmEvent
type DeleteAlarmEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryAlarmView QueryAlarm
type QueryAlarmView struct {
	Inventories []AlarmInventoryView `json:"inventories,omitempty"`
}

// CreateAlarmEventView CreateAlarmEvent
type CreateAlarmEventView struct {
	Inventory AlarmInventoryView `json:"inventory,omitempty"`
}

// AddActionToAlarmEventView AddActionToAlarmEvent
type AddActionToAlarmEventView struct {
	Inventory AlarmInventoryView `json:"inventory,omitempty"`
}

