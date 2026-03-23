// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdatePrimaryStorageParamDetail UpdatePrimaryStorage detail param
type UpdatePrimaryStorageParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Url *string `json:"url,omitempty"`
}

// UpdatePrimaryStorageParam UpdatePrimaryStorage request param
type UpdatePrimaryStorageParam struct {
	BaseParam
	Params UpdatePrimaryStorageParamDetail `json:"updatePrimaryStorage"`
}
// ReconnectPrimaryStorageParamDetail ReconnectPrimaryStorage detail param
type ReconnectPrimaryStorageParamDetail struct {
}

// ReconnectPrimaryStorageParam ReconnectPrimaryStorage request param
type ReconnectPrimaryStorageParam struct {
	BaseParam
	Params ReconnectPrimaryStorageParamDetail `json:"reconnectPrimaryStorage"`
}
// DeletePrimaryStorageParamDetail DeletePrimaryStorage detail param
type DeletePrimaryStorageParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePrimaryStorageParam DeletePrimaryStorage request param
type DeletePrimaryStorageParam struct {
	BaseParam
	Params DeletePrimaryStorageParamDetail `json:"deletePrimaryStorage"`
}
