// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateAlarmParamDetail UpdateAlarm detail param
type UpdateAlarmParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ComparisonOperator *string `json:"comparisonOperator,omitempty"`
	Period *int `json:"period,omitempty"`
	Threshold *float64 `json:"threshold,omitempty"`
	RepeatInterval *int `json:"repeatInterval,omitempty"`
	RepeatCount *int `json:"repeatCount,omitempty"`
	EnableRecovery *bool `json:"enableRecovery,omitempty"`
	EmergencyLevel *string `json:"emergencyLevel,omitempty"`
	Actions []CreateAlarm_ActionParamParam `json:"actions,omitempty"`
}

// UpdateAlarmParam UpdateAlarm request param
type UpdateAlarmParam struct {
	BaseParam
	Params UpdateAlarmParamDetail `json:"updateAlarm"`
}
// DeleteAlarmParamDetail DeleteAlarm detail param
type DeleteAlarmParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAlarmParam DeleteAlarm request param
type DeleteAlarmParam struct {
	BaseParam
	Params DeleteAlarmParamDetail `json:"deleteAlarm"`
}
// CreateAlarmParamDetail CreateAlarm detail param
type CreateAlarmParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ComparisonOperator string `json:"comparisonOperator" validate:"required"`
	Period *int `json:"period,omitempty"`
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	Threshold float64 `json:"threshold" validate:"required"`
	RepeatInterval *int `json:"repeatInterval,omitempty"`
	Labels []LabelParam `json:"labels,omitempty"`
	Actions []CreateAlarm_ActionParamParam `json:"actions,omitempty"`
	RepeatCount *int `json:"repeatCount,omitempty"`
	Type *string `json:"type,omitempty"`
	EnableRecovery *bool `json:"enableRecovery,omitempty"`
	EmergencyLevel *string `json:"emergencyLevel,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAlarmParam CreateAlarm request param
type CreateAlarmParam struct {
	BaseParam
	Params CreateAlarmParamDetail `json:"params"`
}
