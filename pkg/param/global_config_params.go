// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateGlobalConfigParamDetail UpdateGlobalConfig detail param
type UpdateGlobalConfigParamDetail struct {
	Value *string `json:"value,omitempty"`
}

// UpdateGlobalConfigParam UpdateGlobalConfig request param
type UpdateGlobalConfigParam struct {
	BaseParam
	Params UpdateGlobalConfigParamDetail `json:"updateGlobalConfig"`
}
// ResetGlobalConfigParamDetail ResetGlobalConfig detail param
type ResetGlobalConfigParamDetail struct {
}

// ResetGlobalConfigParam ResetGlobalConfig request param
type ResetGlobalConfigParam struct {
	BaseParam
	Params ResetGlobalConfigParamDetail `json:"resetGlobalConfig"`
}
