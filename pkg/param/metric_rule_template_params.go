// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteMetricRuleTemplateParamDetail DeleteMetricRuleTemplate detail param
type DeleteMetricRuleTemplateParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteMetricRuleTemplateParam DeleteMetricRuleTemplate request param
type DeleteMetricRuleTemplateParam struct {
	BaseParam
	Params DeleteMetricRuleTemplateParamDetail `json:"deleteMetricRuleTemplate"`
}
// UpdateMetricRuleTemplateParamDetail UpdateMetricRuleTemplate detail param
type UpdateMetricRuleTemplateParamDetail struct {
	Name string `json:"name,omitempty"`
	ComparisonOperator *string `json:"comparisonOperator,omitempty"`
	Period *int `json:"period,omitempty"`
	Threshold *float64 `json:"threshold,omitempty"`
	RepeatInterval *int `json:"repeatInterval,omitempty"`
	Labels []LabelParam `json:"labels,omitempty"`
	RepeatCount *int `json:"repeatCount,omitempty"`
	EnableRecovery *bool `json:"enableRecovery,omitempty"`
	EmergencyLevel *string `json:"emergencyLevel,omitempty"`
}

// UpdateMetricRuleTemplateParam UpdateMetricRuleTemplate request param
type UpdateMetricRuleTemplateParam struct {
	BaseParam
	Params UpdateMetricRuleTemplateParamDetail `json:"updateMetricRuleTemplate"`
}
// AddMetricRuleTemplateParamDetail AddMetricRuleTemplate detail param
type AddMetricRuleTemplateParamDetail struct {
	Name string `json:"name" validate:"required"`
	ComparisonOperator string `json:"comparisonOperator" validate:"required"`
	Period *int `json:"period,omitempty"`
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	Threshold float64 `json:"threshold" validate:"required"`
	RepeatInterval *int `json:"repeatInterval,omitempty"`
	Labels []LabelParam `json:"labels,omitempty"`
	RepeatCount *int `json:"repeatCount,omitempty"`
	EnableRecovery *bool `json:"enableRecovery,omitempty"`
	EmergencyLevel *string `json:"emergencyLevel,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddMetricRuleTemplateParam AddMetricRuleTemplate request param
type AddMetricRuleTemplateParam struct {
	BaseParam
	Params AddMetricRuleTemplateParamDetail `json:"params"`
}
