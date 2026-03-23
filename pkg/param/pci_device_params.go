// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeletePciDeviceParamDetail DeletePciDevice detail param
type DeletePciDeviceParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePciDeviceParam DeletePciDevice request param
type DeletePciDeviceParam struct {
	BaseParam
	Params DeletePciDeviceParamDetail `json:"deletePciDevice"`
}
// UpdatePciDeviceParamDetail UpdatePciDevice detail param
type UpdatePciDeviceParamDetail struct {
	State *string `json:"state,omitempty"`
	PassThroughState *string `json:"passThroughState,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	MetaData *string `json:"metaData,omitempty"`
}

// UpdatePciDeviceParam UpdatePciDevice request param
type UpdatePciDeviceParam struct {
	BaseParam
	Params UpdatePciDeviceParamDetail `json:"updatePciDevice"`
}
