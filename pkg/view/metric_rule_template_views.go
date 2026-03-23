// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MetricRuleTemplateInventoryView MetricRuleTemplate
type MetricRuleTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	MonitorTemplateUuid string `json:"monitorTemplateUuid,omitempty"`
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	Period int `json:"period,omitempty"`
	RepeatInterval int `json:"repeatInterval,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	MetricName string `json:"metricName,omitempty"`
	Threshold float64 `json:"threshold,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Labels string `json:"labels,omitempty"`
	EnableRecovery bool `json:"enableRecovery,omitempty"`
}

// QueryMetricRuleTemplateView QueryMetricRuleTemplate
type QueryMetricRuleTemplateView struct {
	Inventories []MetricRuleTemplateInventoryView `json:"inventories,omitempty"`
}

// DeleteMetricRuleTemplateEventView DeleteMetricRuleTemplateEvent
type DeleteMetricRuleTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateMetricRuleTemplateEventView UpdateMetricRuleTemplateEvent
type UpdateMetricRuleTemplateEventView struct {
	Inventory MetricRuleTemplateInventoryView `json:"inventory,omitempty"`
}

// AddMetricRuleTemplateEventView AddMetricRuleTemplateEvent
type AddMetricRuleTemplateEventView struct {
	Inventory MetricRuleTemplateInventoryView `json:"inventory,omitempty"`
}

