// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MetricTemplateInventoryView MetricTemplate
type MetricTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	ReceiverUuid string `json:"receiverUuid,omitempty"`
	Template string `json:"template,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	MetricName string `json:"metricName,omitempty"`
	LabelsJsonStr string `json:"labelsJsonStr,omitempty"`
	Description string `json:"description,omitempty"`
}

// DeleteMetricTemplateEventView DeleteMetricTemplateEvent
type DeleteMetricTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateMetricTemplateEventView CreateMetricTemplateEvent
type CreateMetricTemplateEventView struct {
	Inventory MetricTemplateInventoryView `json:"inventory,omitempty"`
}

