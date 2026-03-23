// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateSchedulerJobParamDetail CreateSchedulerJob detail param
type CreateSchedulerJobParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	TargetResourceUuid string `json:"targetResourceUuid" validate:"required"`
	Type string `json:"type" validate:"required"`
	Parameters map[string]string `json:"parameters,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSchedulerJobParam CreateSchedulerJob request param
type CreateSchedulerJobParam struct {
	BaseParam
	Params CreateSchedulerJobParamDetail `json:"params"`
}
// UpdateSchedulerJobParamDetail UpdateSchedulerJob detail param
type UpdateSchedulerJobParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

// UpdateSchedulerJobParam UpdateSchedulerJob request param
type UpdateSchedulerJobParam struct {
	BaseParam
	Params UpdateSchedulerJobParamDetail `json:"updateSchedulerJob"`
}
// DeleteSchedulerJobParamDetail DeleteSchedulerJob detail param
type DeleteSchedulerJobParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteSchedulerJobParam DeleteSchedulerJob request param
type DeleteSchedulerJobParam struct {
	BaseParam
	Params DeleteSchedulerJobParamDetail `json:"deleteSchedulerJob"`
}
