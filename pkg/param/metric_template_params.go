// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteMetricTemplateParamDetail DeleteMetricTemplate detail param
type DeleteMetricTemplateParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteMetricTemplateParam DeleteMetricTemplate request param
type DeleteMetricTemplateParam struct {
	BaseParam
	Params DeleteMetricTemplateParamDetail `json:"deleteMetricTemplate"`
}
// CreateMetricTemplateParamDetail CreateMetricTemplate detail param
type CreateMetricTemplateParamDetail struct {
	ReceiverUuid string `json:"receiverUuid" validate:"required"`
	Template string `json:"template" validate:"required"`
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	LabelsJsonStr *string `json:"labelsJsonStr,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMetricTemplateParam CreateMetricTemplate request param
type CreateMetricTemplateParam struct {
	BaseParam
	Params CreateMetricTemplateParamDetail `json:"params"`
}
