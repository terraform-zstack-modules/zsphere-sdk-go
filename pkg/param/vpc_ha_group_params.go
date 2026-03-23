// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteVpcHaGroupParamDetail DeleteVpcHaGroup detail param
type DeleteVpcHaGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVpcHaGroupParam DeleteVpcHaGroup request param
type DeleteVpcHaGroupParam struct {
	BaseParam
	Params DeleteVpcHaGroupParamDetail `json:"deleteVpcHaGroup"`
}
// CreateVpcHaGroupParamDetail CreateVpcHaGroup detail param
type CreateVpcHaGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	MonitorIps []string `json:"monitorIps,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpcHaGroupParam CreateVpcHaGroup request param
type CreateVpcHaGroupParam struct {
	BaseParam
	Params CreateVpcHaGroupParamDetail `json:"params"`
}
// UpdateVpcHaGroupParamDetail UpdateVpcHaGroup detail param
type UpdateVpcHaGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateVpcHaGroupParam UpdateVpcHaGroup request param
type UpdateVpcHaGroupParam struct {
	BaseParam
	Params UpdateVpcHaGroupParamDetail `json:"updateVpcHaGroup"`
}
