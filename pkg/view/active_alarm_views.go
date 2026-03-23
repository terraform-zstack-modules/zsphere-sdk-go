// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ActiveAlarmInventoryView ActiveAlarm
type ActiveAlarmInventoryView struct {
	BaseInfoView
	BaseTimeView
	TemplateUuid string `json:"templateUuid,omitempty"`
	AlarmUuid string `json:"alarmUuid,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// QueryActiveAlarmView QueryActiveAlarm
type QueryActiveAlarmView struct {
	Inventories []ActiveAlarmInventoryView `json:"inventories,omitempty"`
}

