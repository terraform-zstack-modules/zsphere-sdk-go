// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteL2NetworkParamDetail DeleteL2Network detail param
type DeleteL2NetworkParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteL2NetworkParam DeleteL2Network request param
type DeleteL2NetworkParam struct {
	BaseParam
	Params DeleteL2NetworkParamDetail `json:"deleteL2Network"`
}
// UpdateL2NetworkParamDetail UpdateL2Network detail param
type UpdateL2NetworkParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateL2NetworkParam UpdateL2Network request param
type UpdateL2NetworkParam struct {
	BaseParam
	Params UpdateL2NetworkParamDetail `json:"updateL2Network"`
}
