// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateVniRangeParamDetail UpdateVniRange detail param
type UpdateVniRangeParamDetail struct {
	Name string `json:"name" validate:"required"`
}

// UpdateVniRangeParam UpdateVniRange request param
type UpdateVniRangeParam struct {
	BaseParam
	Params UpdateVniRangeParamDetail `json:"updateVniRange"`
}
// CreateVniRangeParamDetail CreateVniRange detail param
type CreateVniRangeParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	StartVni int `json:"startVni" validate:"required"`
	EndVni int `json:"endVni" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVniRangeParam CreateVniRange request param
type CreateVniRangeParam struct {
	BaseParam
	Params CreateVniRangeParamDetail `json:"params"`
}
// DeleteVniRangeParamDetail DeleteVniRange detail param
type DeleteVniRangeParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVniRangeParam DeleteVniRange request param
type DeleteVniRangeParam struct {
	BaseParam
	Params DeleteVniRangeParamDetail `json:"deleteVniRange"`
}
