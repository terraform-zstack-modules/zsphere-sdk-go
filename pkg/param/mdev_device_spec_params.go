// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateMdevDeviceSpecParamDetail UpdateMdevDeviceSpec detail param
type UpdateMdevDeviceSpecParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
}

// UpdateMdevDeviceSpecParam UpdateMdevDeviceSpec request param
type UpdateMdevDeviceSpecParam struct {
	BaseParam
	Params UpdateMdevDeviceSpecParamDetail `json:"updateMdevDeviceSpec"`
}
