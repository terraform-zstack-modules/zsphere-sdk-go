// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteReservedIpRangeParamDetail DeleteReservedIpRange detail param
type DeleteReservedIpRangeParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteReservedIpRangeParam DeleteReservedIpRange request param
type DeleteReservedIpRangeParam struct {
	BaseParam
	Params DeleteReservedIpRangeParamDetail `json:"deleteReservedIpRange"`
}
// AddReservedIpRangeParamDetail AddReservedIpRange detail param
type AddReservedIpRangeParamDetail struct {
	StartIp string `json:"startIp" validate:"required"`
	EndIp string `json:"endIp" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddReservedIpRangeParam AddReservedIpRange request param
type AddReservedIpRangeParam struct {
	BaseParam
	Params AddReservedIpRangeParamDetail `json:"params"`
}
