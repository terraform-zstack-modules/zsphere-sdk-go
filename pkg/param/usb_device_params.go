// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateUsbDeviceParamDetail UpdateUsbDevice detail param
type UpdateUsbDeviceParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
}

// UpdateUsbDeviceParam UpdateUsbDevice request param
type UpdateUsbDeviceParam struct {
	BaseParam
	Params UpdateUsbDeviceParamDetail `json:"updateUsbDevice"`
}
