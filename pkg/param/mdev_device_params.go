// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateMdevDeviceParamDetail UpdateMdevDevice detail param
type UpdateMdevDeviceParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
}

// UpdateMdevDeviceParam UpdateMdevDevice request param
type UpdateMdevDeviceParam struct {
	BaseParam
	Params UpdateMdevDeviceParamDetail `json:"updateMdevDevice"`
}
// DeleteMdevDeviceParamDetail DeleteMdevDevice detail param
type DeleteMdevDeviceParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteMdevDeviceParam DeleteMdevDevice request param
type DeleteMdevDeviceParam struct {
	BaseParam
	Params DeleteMdevDeviceParamDetail `json:"deleteMdevDevice"`
}
