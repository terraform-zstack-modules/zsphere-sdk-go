// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateVmCdRomParamDetail CreateVmCdRom detail param
type CreateVmCdRomParamDetail struct {
	Name string `json:"name" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	IsoUuid *string `json:"isoUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmCdRomParam CreateVmCdRom request param
type CreateVmCdRomParam struct {
	BaseParam
	Params CreateVmCdRomParamDetail `json:"params"`
}
// UpdateVmCdRomParamDetail UpdateVmCdRom detail param
type UpdateVmCdRomParamDetail struct {
	Description *string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
}

// UpdateVmCdRomParam UpdateVmCdRom request param
type UpdateVmCdRomParam struct {
	BaseParam
	Params UpdateVmCdRomParamDetail `json:"updateVmCdRom"`
}
// DeleteVmCdRomParamDetail DeleteVmCdRom detail param
type DeleteVmCdRomParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVmCdRomParam DeleteVmCdRom request param
type DeleteVmCdRomParam struct {
	BaseParam
	Params DeleteVmCdRomParamDetail `json:"deleteVmCdRom"`
}
