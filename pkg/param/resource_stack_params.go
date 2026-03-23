// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteResourceStackParamDetail DeleteResourceStack detail param
type DeleteResourceStackParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteResourceStackParam DeleteResourceStack request param
type DeleteResourceStackParam struct {
	BaseParam
	Params DeleteResourceStackParamDetail `json:"deleteResourceStack"`
}
// CreateResourceStackParamDetail CreateResourceStack detail param
type CreateResourceStackParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	Rollback *bool `json:"rollback,omitempty"`
	TemplateContent *string `json:"templateContent,omitempty"`
	TemplateUuid *string `json:"templateUuid,omitempty"`
	Parameters *string `json:"parameters,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateResourceStackParam CreateResourceStack request param
type CreateResourceStackParam struct {
	BaseParam
	Params CreateResourceStackParamDetail `json:"params"`
}
// UpdateResourceStackParamDetail UpdateResourceStack detail param
type UpdateResourceStackParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Rollback *bool `json:"rollback,omitempty"`
	TemplateContent *string `json:"templateContent,omitempty"`
	Parameters *string `json:"parameters,omitempty"`
}

// UpdateResourceStackParam UpdateResourceStack request param
type UpdateResourceStackParam struct {
	BaseParam
	Params UpdateResourceStackParamDetail `json:"updateResourceStack"`
}
