// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteVRouterRouteEntryParamDetail DeleteVRouterRouteEntry detail param
type DeleteVRouterRouteEntryParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVRouterRouteEntryParam DeleteVRouterRouteEntry request param
type DeleteVRouterRouteEntryParam struct {
	BaseParam
	Params DeleteVRouterRouteEntryParamDetail `json:"deleteVRouterRouteEntry"`
}
// AddVRouterRouteEntryParamDetail AddVRouterRouteEntry detail param
type AddVRouterRouteEntryParamDetail struct {
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	Destination string `json:"destination" validate:"required"`
	Target *string `json:"target,omitempty"`
	Distance *int `json:"distance,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddVRouterRouteEntryParam AddVRouterRouteEntry request param
type AddVRouterRouteEntryParam struct {
	BaseParam
	Params AddVRouterRouteEntryParamDetail `json:"params"`
}
