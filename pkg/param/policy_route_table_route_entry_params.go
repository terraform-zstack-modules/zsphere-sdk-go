// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeletePolicyRouteTableRouteEntryParamDetail DeletePolicyRouteTableRouteEntry detail param
type DeletePolicyRouteTableRouteEntryParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePolicyRouteTableRouteEntryParam DeletePolicyRouteTableRouteEntry request param
type DeletePolicyRouteTableRouteEntryParam struct {
	BaseParam
	Params DeletePolicyRouteTableRouteEntryParamDetail `json:"deletePolicyRouteTableRouteEntry"`
}
// CreatePolicyRouteTableRouteEntryParamDetail CreatePolicyRouteTableRouteEntry detail param
type CreatePolicyRouteTableRouteEntryParamDetail struct {
	TableUuid string `json:"tableUuid" validate:"required"`
	DestinationCidr string `json:"destinationCidr" validate:"required"`
	NextHopIp string `json:"nextHopIp" validate:"required"`
	Distance *int `json:"distance,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePolicyRouteTableRouteEntryParam CreatePolicyRouteTableRouteEntry request param
type CreatePolicyRouteTableRouteEntryParam struct {
	BaseParam
	Params CreatePolicyRouteTableRouteEntryParamDetail `json:"params"`
}
