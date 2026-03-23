// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MonitorTriggerInventoryView MonitorTrigger
type MonitorTriggerInventoryView struct {
	BaseInfoView
	BaseTimeView
	Expression string `json:"expression,omitempty"`
	RecoveryExpression string `json:"recoveryExpression,omitempty"`
	Description string `json:"description,omitempty"`
	Status string `json:"status,omitempty"`
	State string `json:"state,omitempty"`
	Duration int `json:"duration,omitempty"`
	TargetResourceUuid string `json:"targetResourceUuid,omitempty"`
	LastStatusChangeTime time.Time `json:"lastStatusChangeTime,omitempty"`
}

// QueryMonitorTriggerView QueryMonitorTrigger
type QueryMonitorTriggerView struct {
	Inventories []MonitorTriggerInventoryView `json:"inventories,omitempty"`
}

// QueryMonitorTriggerActionView QueryMonitorTriggerAction
type QueryMonitorTriggerActionView struct {
	Inventories []MonitorTriggerActionInventoryView `json:"inventories,omitempty"`
}

// DeleteMonitorTriggerEventView DeleteMonitorTriggerEvent
type DeleteMonitorTriggerEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateMonitorTriggerEventView UpdateMonitorTriggerEvent
type UpdateMonitorTriggerEventView struct {
	Inventory MonitorTriggerInventoryView `json:"inventory,omitempty"`
}

// CreateMonitorTriggerEventView CreateMonitorTriggerEvent
type CreateMonitorTriggerEventView struct {
	Inventory MonitorTriggerInventoryView `json:"inventory,omitempty"`
}

// DeleteMonitorTriggerActionEventView DeleteMonitorTriggerActionEvent
type DeleteMonitorTriggerActionEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeMonitorTriggerStateEventView ChangeMonitorTriggerStateEvent
type ChangeMonitorTriggerStateEventView struct {
	Inventory MonitorTriggerInventoryView `json:"inventory,omitempty"`
}

