// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateMonitorGroupParamDetail CreateMonitorGroup detail param
type CreateMonitorGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Actions []CreateAlarm_ActionParamParam `json:"actions,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMonitorGroupParam CreateMonitorGroup request param
type CreateMonitorGroupParam struct {
	BaseParam
	Params CreateMonitorGroupParamDetail `json:"params"`
}
// DeleteMonitorGroupParamDetail DeleteMonitorGroup detail param
type DeleteMonitorGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteMonitorGroupParam DeleteMonitorGroup request param
type DeleteMonitorGroupParam struct {
	BaseParam
	Params DeleteMonitorGroupParamDetail `json:"deleteMonitorGroup"`
}
// UpdateMonitorGroupParamDetail UpdateMonitorGroup detail param
type UpdateMonitorGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Actions []CreateAlarm_ActionParamParam `json:"actions,omitempty"`
	StateEvent *string `json:"stateEvent,omitempty"`
}

// UpdateMonitorGroupParam UpdateMonitorGroup request param
type UpdateMonitorGroupParam struct {
	BaseParam
	Params UpdateMonitorGroupParamDetail `json:"updateMonitorGroup"`
}
