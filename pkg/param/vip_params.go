// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateVipParamDetail UpdateVip detail param
type UpdateVipParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateVipParam UpdateVip request param
type UpdateVipParam struct {
	BaseParam
	Params UpdateVipParamDetail `json:"updateVip"`
}
// CreateVipParamDetail CreateVip detail param
type CreateVipParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	AllocatorStrategy *string `json:"allocatorStrategy,omitempty"`
	IpRangeUuid *string `json:"ipRangeUuid,omitempty"`
	RequiredIp *string `json:"requiredIp,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVipParam CreateVip request param
type CreateVipParam struct {
	BaseParam
	Params CreateVipParamDetail `json:"params"`
}
// DeleteVipParamDetail DeleteVip detail param
type DeleteVipParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVipParam DeleteVip request param
type DeleteVipParam struct {
	BaseParam
	Params DeleteVipParamDetail `json:"deleteVip"`
}
