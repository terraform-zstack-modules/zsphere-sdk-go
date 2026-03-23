// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MonitorGroupAlarmInventoryView MonitorGroupAlarm
type MonitorGroupAlarmInventoryView struct {
	BaseInfoView
	BaseTimeView
	GroupUuid string `json:"groupUuid,omitempty"`
	AlarmUuid string `json:"alarmUuid,omitempty"`
	MetricRuleTemplateUuid string `json:"metricRuleTemplateUuid,omitempty"`
}

// QueryMonitorGroupAlarmView QueryMonitorGroupAlarm
type QueryMonitorGroupAlarmView struct {
	Inventories []MonitorGroupAlarmInventoryView `json:"inventories,omitempty"`
}

