// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// GetVirtualRouterSoftwareVersionParamDetail GetVirtualRouterSoftwareVersion detail param
type GetVirtualRouterSoftwareVersionParamDetail struct {
	Uuid string `json:"uuid,omitempty"`
	SoftwareName string `json:"softwareName" validate:"required"`
	NeedUpdate *bool `json:"needUpdate,omitempty"`
}

// GetVirtualRouterSoftwareVersionParam GetVirtualRouterSoftwareVersion request param
type GetVirtualRouterSoftwareVersionParam struct {
	BaseParam
	Params GetVirtualRouterSoftwareVersionParamDetail `json:"getVirtualRouterSoftwareVersion"`
}
// UpdateVirtualRouterSoftwareVersionParamDetail UpdateVirtualRouterSoftwareVersion detail param
type UpdateVirtualRouterSoftwareVersionParamDetail struct {
	SoftwareName string `json:"softwareName" validate:"required"`
	TargetVersion string `json:"targetVersion" validate:"required"`
}

// UpdateVirtualRouterSoftwareVersionParam UpdateVirtualRouterSoftwareVersion request param
type UpdateVirtualRouterSoftwareVersionParam struct {
	BaseParam
	Params UpdateVirtualRouterSoftwareVersionParamDetail `json:"params"`
}
