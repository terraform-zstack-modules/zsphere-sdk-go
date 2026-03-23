// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateXskyBlockVolumeParamDetail UpdateXskyBlockVolume detail param
type UpdateXskyBlockVolumeParamDetail struct {
	BurstTotalBw *int64 `json:"burstTotalBw,omitempty"`
	BurstTotalIops *int64 `json:"burstTotalIops,omitempty"`
	MaxTotalBw *int64 `json:"maxTotalBw,omitempty"`
	MaxTotalIops *int64 `json:"maxTotalIops,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateXskyBlockVolumeParam UpdateXskyBlockVolume request param
type UpdateXskyBlockVolumeParam struct {
	BaseParam
	Params UpdateXskyBlockVolumeParamDetail `json:"updateXskyBlockVolume"`
}
