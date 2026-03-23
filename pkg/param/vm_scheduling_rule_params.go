// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// ValidateVmSchedulingRuleParamDetail ValidateVmSchedulingRule detail param
type ValidateVmSchedulingRuleParamDetail struct {
	VmGroupUuid string `json:"vmGroupUuid" validate:"required"`
	HostGroupUuid *string `json:"hostGroupUuid,omitempty"`
	Rule string `json:"rule" validate:"required"`
	Mode string `json:"mode" validate:"required"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
}

// ValidateVmSchedulingRuleParam ValidateVmSchedulingRule request param
type ValidateVmSchedulingRuleParam struct {
	BaseParam
	Params ValidateVmSchedulingRuleParamDetail `json:"validateVmSchedulingRule"`
}
// UpdateVmSchedulingRuleParamDetail UpdateVmSchedulingRule detail param
type UpdateVmSchedulingRuleParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Mode *string `json:"mode,omitempty"`
}

// UpdateVmSchedulingRuleParam UpdateVmSchedulingRule request param
type UpdateVmSchedulingRuleParam struct {
	BaseParam
	Params UpdateVmSchedulingRuleParamDetail `json:"updateVmSchedulingRule"`
}
// RemoveVmSchedulingRuleParamDetail RemoveVmSchedulingRule detail param
type RemoveVmSchedulingRuleParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveVmSchedulingRuleParam RemoveVmSchedulingRule request param
type RemoveVmSchedulingRuleParam struct {
	BaseParam
	Params RemoveVmSchedulingRuleParamDetail `json:"removeVmSchedulingRule"`
}
// CreateVmSchedulingRuleParamDetail CreateVmSchedulingRule detail param
type CreateVmSchedulingRuleParamDetail struct {
	Rule string `json:"rule" validate:"required"`
	Mode string `json:"mode" validate:"required"`
	VmGroupUuid string `json:"vmGroupUuid" validate:"required"`
	HostGroupUuid *string `json:"hostGroupUuid,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Policy *string `json:"policy,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	SubType *string `json:"subType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmSchedulingRuleParam CreateVmSchedulingRule request param
type CreateVmSchedulingRuleParam struct {
	BaseParam
	Params CreateVmSchedulingRuleParamDetail `json:"params"`
}
