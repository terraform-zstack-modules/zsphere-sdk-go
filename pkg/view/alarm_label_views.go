// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AlarmLabelInventoryView AlarmLabel
type AlarmLabelInventoryView struct {
	BaseInfoView
	BaseTimeView
	Key string `json:"key,omitempty"`
	Operator string `json:"operator,omitempty"`
	Value string `json:"value,omitempty"`
}

// AddLabelToAlarmEventView AddLabelToAlarmEvent
type AddLabelToAlarmEventView struct {
	Inventory AlarmLabelInventoryView `json:"inventory,omitempty"`
}

// UpdateAlarmLabelEventView UpdateAlarmLabelEvent
type UpdateAlarmLabelEventView struct {
	Inventory AlarmLabelInventoryView `json:"inventory,omitempty"`
}

