// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateVmSchedulingRuleGroupParamDetail UpdateVmSchedulingRuleGroup detail param
type UpdateVmSchedulingRuleGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateVmSchedulingRuleGroupParam UpdateVmSchedulingRuleGroup request param
type UpdateVmSchedulingRuleGroupParam struct {
	BaseParam
	Params UpdateVmSchedulingRuleGroupParamDetail `json:"updateVmSchedulingRuleGroup"`
}
// DeleteVmSchedulingRuleGroupParamDetail DeleteVmSchedulingRuleGroup detail param
type DeleteVmSchedulingRuleGroupParamDetail struct {
}

// DeleteVmSchedulingRuleGroupParam DeleteVmSchedulingRuleGroup request param
type DeleteVmSchedulingRuleGroupParam struct {
	BaseParam
	Params DeleteVmSchedulingRuleGroupParamDetail `json:"deleteVmSchedulingRuleGroup"`
}
// CreateVmSchedulingRuleGroupParamDetail CreateVmSchedulingRuleGroup detail param
type CreateVmSchedulingRuleGroupParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmSchedulingRuleGroupParam CreateVmSchedulingRuleGroup request param
type CreateVmSchedulingRuleGroupParam struct {
	BaseParam
	Params CreateVmSchedulingRuleGroupParamDetail `json:"params"`
}
