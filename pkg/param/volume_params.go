// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateVolumeParamDetail UpdateVolume detail param
type UpdateVolumeParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateVolumeParam UpdateVolume request param
type UpdateVolumeParam struct {
	BaseParam
	Params UpdateVolumeParamDetail `json:"updateVolume"`
}
