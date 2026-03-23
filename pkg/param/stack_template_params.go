// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteStackTemplateParamDetail DeleteStackTemplate detail param
type DeleteStackTemplateParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteStackTemplateParam DeleteStackTemplate request param
type DeleteStackTemplateParam struct {
	BaseParam
	Params DeleteStackTemplateParamDetail `json:"deleteStackTemplate"`
}
// UpdateStackTemplateParamDetail UpdateStackTemplate detail param
type UpdateStackTemplateParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	State *bool `json:"state,omitempty"`
	TemplateContent *string `json:"templateContent,omitempty"`
}

// UpdateStackTemplateParam UpdateStackTemplate request param
type UpdateStackTemplateParam struct {
	BaseParam
	Params UpdateStackTemplateParamDetail `json:"updateStackTemplate"`
}
// AddStackTemplateParamDetail AddStackTemplate detail param
type AddStackTemplateParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	TemplateContent *string `json:"templateContent,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddStackTemplateParam AddStackTemplate request param
type AddStackTemplateParam struct {
	BaseParam
	Params AddStackTemplateParamDetail `json:"params"`
}
