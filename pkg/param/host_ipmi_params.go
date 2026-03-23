// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateHostIpmiParamDetail UpdateHostIpmi detail param
type UpdateHostIpmiParamDetail struct {
	IpmiAddress *string `json:"ipmiAddress,omitempty"`
	IpmiUsername *string `json:"ipmiUsername,omitempty"`
	IpmiPassword *string `json:"ipmiPassword,omitempty"`
	IpmiPort *int `json:"ipmiPort,omitempty"`
}

// UpdateHostIpmiParam UpdateHostIpmi request param
type UpdateHostIpmiParam struct {
	BaseParam
	Params UpdateHostIpmiParamDetail `json:"updateHostIpmi"`
}
