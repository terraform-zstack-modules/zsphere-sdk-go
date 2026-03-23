// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MonitorGroupTemplateRefInventoryView MonitorGroupTemplateRef
type MonitorGroupTemplateRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	TemplateUuid string `json:"templateUuid,omitempty"`
	GroupUuid string `json:"groupUuid,omitempty"`
	IsApplied bool `json:"isApplied,omitempty"`
}

// ApplyMonitorTemplateToMonitorGroupEventView ApplyMonitorTemplateToMonitorGroupEvent
type ApplyMonitorTemplateToMonitorGroupEventView struct {
	Inventory MonitorGroupTemplateRefInventoryView `json:"inventory,omitempty"`
}

// QueryMonitorGroupTemplateRefView QueryMonitorGroupTemplateRef
type QueryMonitorGroupTemplateRefView struct {
	Inventories []MonitorGroupTemplateRefInventoryView `json:"inventories,omitempty"`
}

