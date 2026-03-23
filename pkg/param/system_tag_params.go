// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateSystemTagParamDetail UpdateSystemTag detail param
type UpdateSystemTagParamDetail struct {
	Tag string `json:"tag" validate:"required"`
}

// UpdateSystemTagParam UpdateSystemTag request param
type UpdateSystemTagParam struct {
	BaseParam
	Params UpdateSystemTagParamDetail `json:"updateSystemTag"`
}
// CreateSystemTagParamDetail CreateSystemTag detail param
type CreateSystemTagParamDetail struct {
	ResourceType string `json:"resourceType" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	Tag string `json:"tag" validate:"required"`
}

// CreateSystemTagParam CreateSystemTag request param
type CreateSystemTagParam struct {
	BaseParam
	Params CreateSystemTagParamDetail `json:"params"`
}
