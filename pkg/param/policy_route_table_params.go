// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreatePolicyRouteTableParamDetail CreatePolicyRouteTable detail param
type CreatePolicyRouteTableParamDetail struct {
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	Number int `json:"number" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePolicyRouteTableParam CreatePolicyRouteTable request param
type CreatePolicyRouteTableParam struct {
	BaseParam
	Params CreatePolicyRouteTableParamDetail `json:"params"`
}
// DeletePolicyRouteTableParamDetail DeletePolicyRouteTable detail param
type DeletePolicyRouteTableParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePolicyRouteTableParam DeletePolicyRouteTable request param
type DeletePolicyRouteTableParam struct {
	BaseParam
	Params DeletePolicyRouteTableParamDetail `json:"deletePolicyRouteTable"`
}
