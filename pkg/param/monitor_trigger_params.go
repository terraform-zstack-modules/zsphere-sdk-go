// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteMonitorTriggerParamDetail DeleteMonitorTrigger detail param
type DeleteMonitorTriggerParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteMonitorTriggerParam DeleteMonitorTrigger request param
type DeleteMonitorTriggerParam struct {
	BaseParam
	Params DeleteMonitorTriggerParamDetail `json:"deleteMonitorTrigger"`
}
// UpdateMonitorTriggerParamDetail UpdateMonitorTrigger detail param
type UpdateMonitorTriggerParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Expression *string `json:"expression,omitempty"`
	Duration *int `json:"duration,omitempty"`
}

// UpdateMonitorTriggerParam UpdateMonitorTrigger request param
type UpdateMonitorTriggerParam struct {
	BaseParam
	Params UpdateMonitorTriggerParamDetail `json:"updateMonitorTrigger"`
}
// CreateMonitorTriggerParamDetail CreateMonitorTrigger detail param
type CreateMonitorTriggerParamDetail struct {
	Name string `json:"name" validate:"required"`
	Expression string `json:"expression" validate:"required"`
	Duration int `json:"duration" validate:"required"`
	RecoveryExpression *string `json:"recoveryExpression,omitempty"`
	Description *string `json:"description,omitempty"`
	TargetResourceUuid string `json:"targetResourceUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMonitorTriggerParam CreateMonitorTrigger request param
type CreateMonitorTriggerParam struct {
	BaseParam
	Params CreateMonitorTriggerParamDetail `json:"params"`
}
// DeleteMonitorTriggerActionParamDetail DeleteMonitorTriggerAction detail param
type DeleteMonitorTriggerActionParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteMonitorTriggerActionParam DeleteMonitorTriggerAction request param
type DeleteMonitorTriggerActionParam struct {
	BaseParam
	Params DeleteMonitorTriggerActionParamDetail `json:"deleteMonitorTriggerAction"`
}
