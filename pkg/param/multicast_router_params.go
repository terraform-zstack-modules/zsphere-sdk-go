// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateMulticastRouterParamDetail CreateMulticastRouter detail param
type CreateMulticastRouterParamDetail struct {
	VpcRouterVmUuid string `json:"vpcRouterVmUuid" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMulticastRouterParam CreateMulticastRouter request param
type CreateMulticastRouterParam struct {
	BaseParam
	Params CreateMulticastRouterParamDetail `json:"params"`
}
// DeleteMulticastRouterParamDetail DeleteMulticastRouter detail param
type DeleteMulticastRouterParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteMulticastRouterParam DeleteMulticastRouter request param
type DeleteMulticastRouterParam struct {
	BaseParam
	Params DeleteMulticastRouterParamDetail `json:"deleteMulticastRouter"`
}
