// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateCbtTaskParamDetail CreateCbtTask detail param
type CreateCbtTaskParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateCbtTaskParam CreateCbtTask request param
type CreateCbtTaskParam struct {
	BaseParam
	Params CreateCbtTaskParamDetail `json:"params"`
}
// DeleteCbtTaskParamDetail DeleteCbtTask detail param
type DeleteCbtTaskParamDetail struct {
	Force *bool `json:"force,omitempty"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteCbtTaskParam DeleteCbtTask request param
type DeleteCbtTaskParam struct {
	BaseParam
	Params DeleteCbtTaskParamDetail `json:"deleteCbtTask"`
}
