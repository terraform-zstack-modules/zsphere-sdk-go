// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// GetVmDnsParamDetail GetVmDns detail param
type GetVmDnsParamDetail struct {
	VmNicUuid *string `json:"vmNicUuid,omitempty"`
	IpVersion *int `json:"ipVersion,omitempty"`
}

// GetVmDnsParam GetVmDns request param
type GetVmDnsParam struct {
	BaseParam
	Params GetVmDnsParamDetail `json:"getVmDns"`
}
// SetVmDnsParamDetail SetVmDns detail param
type SetVmDnsParamDetail struct {
	VmNicUuid *string `json:"vmNicUuid,omitempty"`
	DnsList []string `json:"dnsList" validate:"required"`
	IpVersion *int `json:"ipVersion,omitempty"`
}

// SetVmDnsParam SetVmDns request param
type SetVmDnsParam struct {
	BaseParam
	Params SetVmDnsParamDetail `json:"setVmDns"`
}
