// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateBaremetalBondingParamDetail CreateBaremetalBonding detail param
type CreateBaremetalBondingParamDetail struct {
	ChassisUuid string `json:"chassisUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Mode int `json:"mode" validate:"required"`
	Slaves string `json:"slaves" validate:"required"`
	Opts *string `json:"opts,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBaremetalBondingParam CreateBaremetalBonding request param
type CreateBaremetalBondingParam struct {
	BaseParam
	Params CreateBaremetalBondingParamDetail `json:"params"`
}
