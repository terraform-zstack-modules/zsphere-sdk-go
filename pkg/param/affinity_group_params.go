// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateAffinityGroupParamDetail UpdateAffinityGroup detail param
type UpdateAffinityGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateAffinityGroupParam UpdateAffinityGroup request param
type UpdateAffinityGroupParam struct {
	BaseParam
	Params UpdateAffinityGroupParamDetail `json:"updateAffinityGroup"`
}
// DeleteAffinityGroupParamDetail DeleteAffinityGroup detail param
type DeleteAffinityGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAffinityGroupParam DeleteAffinityGroup request param
type DeleteAffinityGroupParam struct {
	BaseParam
	Params DeleteAffinityGroupParamDetail `json:"deleteAffinityGroup"`
}
// CreateAffinityGroupParamDetail CreateAffinityGroup detail param
type CreateAffinityGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Policy *string `json:"policy,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	SubType *string `json:"subType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAffinityGroupParam CreateAffinityGroup request param
type CreateAffinityGroupParam struct {
	BaseParam
	Params CreateAffinityGroupParamDetail `json:"params"`
}
