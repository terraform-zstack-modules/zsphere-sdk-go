// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateAlarmLabelParamDetail UpdateAlarmLabel detail param
type UpdateAlarmLabelParamDetail struct {
	Key string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
	Operator string `json:"operator" validate:"required"`
}

// UpdateAlarmLabelParam UpdateAlarmLabel request param
type UpdateAlarmLabelParam struct {
	BaseParam
	Params UpdateAlarmLabelParamDetail `json:"updateAlarmLabel"`
}
