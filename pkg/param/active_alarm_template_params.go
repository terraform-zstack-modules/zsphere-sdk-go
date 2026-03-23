// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateActiveAlarmTemplateParamDetail UpdateActiveAlarmTemplate detail param
type UpdateActiveAlarmTemplateParamDetail struct {
	AlarmName *string `json:"alarmName,omitempty"`
	ComparisonOperator *string `json:"comparisonOperator,omitempty"`
	Period *int `json:"period,omitempty"`
	Threshold *float64 `json:"threshold,omitempty"`
	RepeatInterval *int `json:"repeatInterval,omitempty"`
	RepeatCount *int `json:"repeatCount,omitempty"`
	EmergencyLevel *string `json:"emergencyLevel,omitempty"`
	Labels *string `json:"labels,omitempty"`
}

// UpdateActiveAlarmTemplateParam UpdateActiveAlarmTemplate request param
type UpdateActiveAlarmTemplateParam struct {
	BaseParam
	Params UpdateActiveAlarmTemplateParamDetail `json:"updateActiveAlarmTemplate"`
}
