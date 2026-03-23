// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddZBoxParamDetail AddZBox detail param
type AddZBoxParamDetail struct {
	UsbDeviceUuid string `json:"usbDeviceUuid" validate:"required"`
	Name string `json:"name,omitempty"`
	SkipFormat *bool `json:"skipFormat,omitempty"`
}

// AddZBoxParam AddZBox request param
type AddZBoxParam struct {
	BaseParam
	Params AddZBoxParamDetail `json:"params"`
}
