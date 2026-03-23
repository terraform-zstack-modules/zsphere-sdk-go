// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// GetVRouterRouteTableParamDetail GetVRouterRouteTable detail param
type GetVRouterRouteTableParamDetail struct {
}

// GetVRouterRouteTableParam GetVRouterRouteTable request param
type GetVRouterRouteTableParam struct {
	BaseParam
	Params GetVRouterRouteTableParamDetail `json:"getVRouterRouteTable"`
}
// CreateVRouterRouteTableParamDetail CreateVRouterRouteTable detail param
type CreateVRouterRouteTableParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVRouterRouteTableParam CreateVRouterRouteTable request param
type CreateVRouterRouteTableParam struct {
	BaseParam
	Params CreateVRouterRouteTableParamDetail `json:"params"`
}
// UpdateVRouterRouteTableParamDetail UpdateVRouterRouteTable detail param
type UpdateVRouterRouteTableParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateVRouterRouteTableParam UpdateVRouterRouteTable request param
type UpdateVRouterRouteTableParam struct {
	BaseParam
	Params UpdateVRouterRouteTableParamDetail `json:"updateVRouterRouteTable"`
}
// DeleteVRouterRouteTableParamDetail DeleteVRouterRouteTable detail param
type DeleteVRouterRouteTableParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVRouterRouteTableParam DeleteVRouterRouteTable request param
type DeleteVRouterRouteTableParam struct {
	BaseParam
	Params DeleteVRouterRouteTableParamDetail `json:"deleteVRouterRouteTable"`
}
