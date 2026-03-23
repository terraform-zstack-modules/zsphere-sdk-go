// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CloneMonitorTemplateParamDetail CloneMonitorTemplate detail param
type CloneMonitorTemplateParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CloneMonitorTemplateParam CloneMonitorTemplate request param
type CloneMonitorTemplateParam struct {
	BaseParam
	Params CloneMonitorTemplateParamDetail `json:"params"`
}
// UpdateMonitorTemplateParamDetail UpdateMonitorTemplate detail param
type UpdateMonitorTemplateParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateMonitorTemplateParam UpdateMonitorTemplate request param
type UpdateMonitorTemplateParam struct {
	BaseParam
	Params UpdateMonitorTemplateParamDetail `json:"updateMonitorTemplate"`
}
// CreateMonitorTemplateParamDetail CreateMonitorTemplate detail param
type CreateMonitorTemplateParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMonitorTemplateParam CreateMonitorTemplate request param
type CreateMonitorTemplateParam struct {
	BaseParam
	Params CreateMonitorTemplateParamDetail `json:"params"`
}
// DeleteMonitorTemplateParamDetail DeleteMonitorTemplate detail param
type DeleteMonitorTemplateParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteMonitorTemplateParam DeleteMonitorTemplate request param
type DeleteMonitorTemplateParam struct {
	BaseParam
	Params DeleteMonitorTemplateParamDetail `json:"deleteMonitorTemplate"`
}
