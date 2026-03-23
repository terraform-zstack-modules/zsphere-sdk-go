// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MonitorTemplateInventoryView MonitorTemplate
type MonitorTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	MonitorGroupTemplateRefs []MonitorGroupTemplateRefVOView `json:"monitorGroupTemplateRefs,omitempty"`
}

// CloneMonitorTemplateEventView CloneMonitorTemplateEvent
type CloneMonitorTemplateEventView struct {
	Inventory MonitorTemplateInventoryView `json:"inventory,omitempty"`
}

// UpdateMonitorTemplateEventView UpdateMonitorTemplateEvent
type UpdateMonitorTemplateEventView struct {
	Inventory MonitorTemplateInventoryView `json:"inventory,omitempty"`
}

// QueryMonitorTemplateView QueryMonitorTemplate
type QueryMonitorTemplateView struct {
	Inventories []MonitorTemplateInventoryView `json:"inventories,omitempty"`
}

// CreateMonitorTemplateEventView CreateMonitorTemplateEvent
type CreateMonitorTemplateEventView struct {
	Inventory MonitorTemplateInventoryView `json:"inventory,omitempty"`
}

// DeleteMonitorTemplateEventView DeleteMonitorTemplateEvent
type DeleteMonitorTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

