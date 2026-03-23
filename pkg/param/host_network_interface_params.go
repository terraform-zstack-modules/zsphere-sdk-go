// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateHostNetworkInterfaceParamDetail UpdateHostNetworkInterface detail param
type UpdateHostNetworkInterfaceParamDetail struct {
	Description *string `json:"description" validate:"required"`
}

// UpdateHostNetworkInterfaceParam UpdateHostNetworkInterface request param
type UpdateHostNetworkInterfaceParam struct {
	BaseParam
	Params UpdateHostNetworkInterfaceParamDetail `json:"params"`
}
// LocateHostNetworkInterfaceParamDetail LocateHostNetworkInterface detail param
type LocateHostNetworkInterfaceParamDetail struct {
	NetworkInterfaceName string `json:"networkInterfaceName" validate:"required"`
	Interval *int64 `json:"interval,omitempty"`
}

// LocateHostNetworkInterfaceParam LocateHostNetworkInterface request param
type LocateHostNetworkInterfaceParam struct {
	BaseParam
	Params LocateHostNetworkInterfaceParamDetail `json:"locateHostNetworkInterface"`
}
