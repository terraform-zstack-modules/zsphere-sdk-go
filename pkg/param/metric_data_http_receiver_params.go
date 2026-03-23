// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteMetricDataHttpReceiverParamDetail DeleteMetricDataHttpReceiver detail param
type DeleteMetricDataHttpReceiverParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteMetricDataHttpReceiverParam DeleteMetricDataHttpReceiver request param
type DeleteMetricDataHttpReceiverParam struct {
	BaseParam
	Params DeleteMetricDataHttpReceiverParamDetail `json:"deleteMetricDataHttpReceiver"`
}
// CreateMetricDataHttpReceiverParamDetail CreateMetricDataHttpReceiver detail param
type CreateMetricDataHttpReceiverParamDetail struct {
	Name string `json:"name" validate:"required"`
	Url string `json:"url" validate:"required"`
	Description *string `json:"description,omitempty"`
	DefaultEnable *bool `json:"defaultEnable,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMetricDataHttpReceiverParam CreateMetricDataHttpReceiver request param
type CreateMetricDataHttpReceiverParam struct {
	BaseParam
	Params CreateMetricDataHttpReceiverParamDetail `json:"params"`
}
