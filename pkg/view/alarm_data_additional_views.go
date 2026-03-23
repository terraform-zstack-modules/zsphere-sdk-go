// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AlarmDataView AlarmData
type AlarmDataView struct {
	AlarmUuid string `json:"alarmUuid,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	MetricName string `json:"metricName,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	AlarmStatus string `json:"alarmStatus,omitempty"`
	AlarmName string `json:"alarmName,omitempty"`
	Threshold float64 `json:"threshold,omitempty"`
	Period int `json:"period,omitempty"`
	Labels string `json:"labels,omitempty"`
	MetricValue float64 `json:"metricValue,omitempty"`
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	ReadStatus string `json:"readStatus,omitempty"`
	DataUuid string `json:"dataUuid,omitempty"`
	Context string `json:"context,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Time int64 `json:"time,omitempty"`
}

