// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdatePciDeviceSpecParamDetail UpdatePciDeviceSpec detail param
type UpdatePciDeviceSpecParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	RomContent *string `json:"romContent,omitempty"`
	RomVersion *string `json:"romVersion,omitempty"`
	AbandonSpecRom *bool `json:"abandonSpecRom,omitempty"`
	State *string `json:"state,omitempty"`
}

// UpdatePciDeviceSpecParam UpdatePciDeviceSpec request param
type UpdatePciDeviceSpecParam struct {
	BaseParam
	Params UpdatePciDeviceSpecParamDetail `json:"updatePciDeviceSpec"`
}
