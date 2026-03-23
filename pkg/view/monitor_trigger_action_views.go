// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MonitorTriggerActionInventoryView MonitorTriggerAction
type MonitorTriggerActionInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	TriggerUuids []string `json:"triggerUuids,omitempty"`
}

// UpdateMonitorTriggerActionEventView UpdateMonitorTriggerActionEvent
type UpdateMonitorTriggerActionEventView struct {
	Inventory MonitorTriggerActionInventoryView `json:"inventory,omitempty"`
}

// CreateMonitorTriggerActionEventView CreateMonitorTriggerActionEvent
type CreateMonitorTriggerActionEventView struct {
	Inventory MonitorTriggerActionInventoryView `json:"inventory,omitempty"`
}

// ChangeMonitorTriggerActionStateEventView ChangeMonitorTriggerActionStateEvent
type ChangeMonitorTriggerActionStateEventView struct {
	Inventory MonitorTriggerActionInventoryView `json:"inventory,omitempty"`
}

