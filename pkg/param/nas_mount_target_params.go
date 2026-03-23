// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateNasMountTargetParamDetail UpdateNasMountTarget detail param
type UpdateNasMountTargetParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateNasMountTargetParam UpdateNasMountTarget request param
type UpdateNasMountTargetParam struct {
	BaseParam
	Params UpdateNasMountTargetParamDetail `json:"updateNasMountTarget"`
}
// DeleteNasMountTargetParamDetail DeleteNasMountTarget detail param
type DeleteNasMountTargetParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteNasMountTargetParam DeleteNasMountTarget request param
type DeleteNasMountTargetParam struct {
	BaseParam
	Params DeleteNasMountTargetParamDetail `json:"deleteNasMountTarget"`
}
