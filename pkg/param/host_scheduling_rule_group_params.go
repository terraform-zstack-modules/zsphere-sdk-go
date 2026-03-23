// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateHostSchedulingRuleGroupParamDetail UpdateHostSchedulingRuleGroup detail param
type UpdateHostSchedulingRuleGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateHostSchedulingRuleGroupParam UpdateHostSchedulingRuleGroup request param
type UpdateHostSchedulingRuleGroupParam struct {
	BaseParam
	Params UpdateHostSchedulingRuleGroupParamDetail `json:"updateHostSchedulingRuleGroup"`
}
// CreateHostSchedulingRuleGroupParamDetail CreateHostSchedulingRuleGroup detail param
type CreateHostSchedulingRuleGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateHostSchedulingRuleGroupParam CreateHostSchedulingRuleGroup request param
type CreateHostSchedulingRuleGroupParam struct {
	BaseParam
	Params CreateHostSchedulingRuleGroupParamDetail `json:"params"`
}
// DeleteHostSchedulingRuleGroupParamDetail DeleteHostSchedulingRuleGroup detail param
type DeleteHostSchedulingRuleGroupParamDetail struct {
}

// DeleteHostSchedulingRuleGroupParam DeleteHostSchedulingRuleGroup request param
type DeleteHostSchedulingRuleGroupParam struct {
	BaseParam
	Params DeleteHostSchedulingRuleGroupParamDetail `json:"deleteHostSchedulingRuleGroup"`
}
