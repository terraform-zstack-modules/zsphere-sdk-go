// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteResourceAttributeKeyParamDetail DeleteResourceAttributeKey detail param
type DeleteResourceAttributeKeyParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteResourceAttributeKeyParam DeleteResourceAttributeKey request param
type DeleteResourceAttributeKeyParam struct {
	BaseParam
	Params DeleteResourceAttributeKeyParamDetail `json:"deleteResourceAttributeKey"`
}
// UpdateResourceAttributeKeyParamDetail UpdateResourceAttributeKey detail param
type UpdateResourceAttributeKeyParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceTypes []string `json:"resourceTypes,omitempty"`
	CreateConstraints []ResourceAttributeConstraintParamParam `json:"createConstraints,omitempty"`
	UpdateConstraints []ResourceAttributeConstraintParamParam `json:"updateConstraints,omitempty"`
	DeleteConstraintIds []int64 `json:"deleteConstraintIds,omitempty"`
}

// UpdateResourceAttributeKeyParam UpdateResourceAttributeKey request param
type UpdateResourceAttributeKeyParam struct {
	BaseParam
	Params UpdateResourceAttributeKeyParamDetail `json:"updateResourceAttributeKey"`
}
// CreateResourceAttributeKeyParamDetail CreateResourceAttributeKey detail param
type CreateResourceAttributeKeyParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceTypes []string `json:"resourceTypes,omitempty"`
	Constraints []ResourceAttributeConstraintParamParam `json:"constraints,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateResourceAttributeKeyParam CreateResourceAttributeKey request param
type CreateResourceAttributeKeyParam struct {
	BaseParam
	Params CreateResourceAttributeKeyParamDetail `json:"params"`
}
