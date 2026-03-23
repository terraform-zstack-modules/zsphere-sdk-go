// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteSNSApplicationPlatformParamDetail DeleteSNSApplicationPlatform detail param
type DeleteSNSApplicationPlatformParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteSNSApplicationPlatformParam DeleteSNSApplicationPlatform request param
type DeleteSNSApplicationPlatformParam struct {
	BaseParam
	Params DeleteSNSApplicationPlatformParamDetail `json:"deleteSNSApplicationPlatform"`
}
// UpdateSNSApplicationPlatformParamDetail UpdateSNSApplicationPlatform detail param
type UpdateSNSApplicationPlatformParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateSNSApplicationPlatformParam UpdateSNSApplicationPlatform request param
type UpdateSNSApplicationPlatformParam struct {
	BaseParam
	Params UpdateSNSApplicationPlatformParamDetail `json:"updateSNSApplicationPlatform"`
}
