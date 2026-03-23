// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateSharedBlockParamDetail UpdateSharedBlock detail param
type UpdateSharedBlockParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	DiskUuid *string `json:"diskUuid,omitempty"`
}

// UpdateSharedBlockParam UpdateSharedBlock request param
type UpdateSharedBlockParam struct {
	BaseParam
	Params UpdateSharedBlockParamDetail `json:"updateSharedBlock"`
}
