// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateLogConfigurationParamDetail UpdateLogConfiguration detail param
type UpdateLogConfigurationParamDetail struct {
	ConfigId int64 `json:"configId" validate:"required"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateLogConfigurationParam UpdateLogConfiguration request param
type UpdateLogConfigurationParam struct {
	BaseParam
	Params UpdateLogConfigurationParamDetail `json:"updateLogConfiguration"`
}
// DeleteLogConfigurationParamDetail DeleteLogConfiguration detail param
type DeleteLogConfigurationParamDetail struct {
	ConfigId int64 `json:"configId" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteLogConfigurationParam DeleteLogConfiguration request param
type DeleteLogConfigurationParam struct {
	BaseParam
	Params DeleteLogConfigurationParamDetail `json:"deleteLogConfiguration"`
}
// AddLogConfigurationParamDetail AddLogConfiguration detail param
type AddLogConfigurationParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	Configuration string `json:"configuration" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddLogConfigurationParam AddLogConfiguration request param
type AddLogConfigurationParam struct {
	BaseParam
	Params AddLogConfigurationParamDetail `json:"params"`
}
// GetLogConfigurationParamDetail GetLogConfiguration detail param
type GetLogConfigurationParamDetail struct {
}

// GetLogConfigurationParam GetLogConfiguration request param
type GetLogConfigurationParam struct {
	BaseParam
	Params GetLogConfigurationParamDetail `json:"getLogConfiguration"`
}
