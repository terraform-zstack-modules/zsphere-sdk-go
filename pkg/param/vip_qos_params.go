// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteVipQosParamDetail DeleteVipQos detail param
type DeleteVipQosParamDetail struct {
	Port *int `json:"port,omitempty"`
}

// DeleteVipQosParam DeleteVipQos request param
type DeleteVipQosParam struct {
	BaseParam
	Params DeleteVipQosParamDetail `json:"deleteVipQos"`
}
// GetVipQosParamDetail GetVipQos detail param
type GetVipQosParamDetail struct {
}

// GetVipQosParam GetVipQos request param
type GetVipQosParam struct {
	BaseParam
	Params GetVipQosParamDetail `json:"getVipQos"`
}
// SetVipQosParamDetail SetVipQos detail param
type SetVipQosParamDetail struct {
	Port *int `json:"port,omitempty"`
	OutboundBandwidth *int64 `json:"outboundBandwidth,omitempty"`
	InboundBandwidth *int64 `json:"inboundBandwidth,omitempty"`
}

// SetVipQosParam SetVipQos request param
type SetVipQosParam struct {
	BaseParam
	Params SetVipQosParamDetail `json:"setVipQos"`
}
