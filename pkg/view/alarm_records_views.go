// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AlarmRecordsInventoryView AlarmRecords
type AlarmRecordsInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	CreateTime int64 `json:"createTime,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	AlarmName string `json:"alarmName,omitempty"`
	AlarmStatus string `json:"alarmStatus,omitempty"`
	AlarmUuid string `json:"alarmUuid,omitempty"`
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	Context string `json:"context,omitempty"`
	DataUuid string `json:"dataUuid,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Labels string `json:"labels,omitempty"`
	MetricName string `json:"metricName,omitempty"`
	MetricValue float64 `json:"metricValue,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Period int `json:"period,omitempty"`
	ReadStatus bool `json:"readStatus,omitempty"`
	OperatorAccountUuid string `json:"operatorAccountUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	Threshold float64 `json:"threshold,omitempty"`
}

// QueryAlarmRecordView QueryAlarmRecord
type QueryAlarmRecordView struct {
	Inventories []AlarmRecordsInventoryView `json:"inventories,omitempty"`
}

