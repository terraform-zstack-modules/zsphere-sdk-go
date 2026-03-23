// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateHostParamDetail UpdateHost detail param
type UpdateHostParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ManagementIp *string `json:"managementIp,omitempty"`
}

// UpdateHostParam UpdateHost request param
type UpdateHostParam struct {
	BaseParam
	Params UpdateHostParamDetail `json:"updateHost"`
}
// ReconnectHostParamDetail ReconnectHost detail param
type ReconnectHostParamDetail struct {
}

// ReconnectHostParam ReconnectHost request param
type ReconnectHostParam struct {
	BaseParam
	Params ReconnectHostParamDetail `json:"reconnectHost"`
}
// DeleteHostParamDetail DeleteHost detail param
type DeleteHostParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteHostParam DeleteHost request param
type DeleteHostParam struct {
	BaseParam
	Params DeleteHostParamDetail `json:"deleteHost"`
}
