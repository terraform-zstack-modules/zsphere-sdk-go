// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateScsiLunParamDetail UpdateScsiLun detail param
type UpdateScsiLunParamDetail struct {
	Name string `json:"name,omitempty"`
	State *string `json:"state,omitempty"`
}

// UpdateScsiLunParam UpdateScsiLun request param
type UpdateScsiLunParam struct {
	BaseParam
	Params UpdateScsiLunParamDetail `json:"updateScsiLun"`
}
