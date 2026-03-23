// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeletePreconfigurationTemplateParamDetail DeletePreconfigurationTemplate detail param
type DeletePreconfigurationTemplateParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePreconfigurationTemplateParam DeletePreconfigurationTemplate request param
type DeletePreconfigurationTemplateParam struct {
	BaseParam
	Params DeletePreconfigurationTemplateParamDetail `json:"deletePreconfigurationTemplate"`
}
// UpdatePreconfigurationTemplateParamDetail UpdatePreconfigurationTemplate detail param
type UpdatePreconfigurationTemplateParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Distribution *string `json:"distribution,omitempty"`
	Type *string `json:"type,omitempty"`
	Content *string `json:"content,omitempty"`
}

// UpdatePreconfigurationTemplateParam UpdatePreconfigurationTemplate request param
type UpdatePreconfigurationTemplateParam struct {
	BaseParam
	Params UpdatePreconfigurationTemplateParamDetail `json:"updatePreconfigurationTemplate"`
}
// AddPreconfigurationTemplateParamDetail AddPreconfigurationTemplate detail param
type AddPreconfigurationTemplateParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Distribution string `json:"distribution" validate:"required"`
	Type string `json:"type" validate:"required"`
	Content string `json:"content" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddPreconfigurationTemplateParam AddPreconfigurationTemplate request param
type AddPreconfigurationTemplateParam struct {
	BaseParam
	Params AddPreconfigurationTemplateParamDetail `json:"params"`
}
