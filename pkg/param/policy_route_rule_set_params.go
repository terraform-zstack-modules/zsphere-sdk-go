// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreatePolicyRouteRuleSetParamDetail CreatePolicyRouteRuleSet detail param
type CreatePolicyRouteRuleSetParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	Type *string `json:"type,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePolicyRouteRuleSetParam CreatePolicyRouteRuleSet request param
type CreatePolicyRouteRuleSetParam struct {
	BaseParam
	Params CreatePolicyRouteRuleSetParamDetail `json:"params"`
}
// UpdatePolicyRouteRuleSetParamDetail UpdatePolicyRouteRuleSet detail param
type UpdatePolicyRouteRuleSetParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdatePolicyRouteRuleSetParam UpdatePolicyRouteRuleSet request param
type UpdatePolicyRouteRuleSetParam struct {
	BaseParam
	Params UpdatePolicyRouteRuleSetParamDetail `json:"updatePolicyRouteRuleSet"`
}
// DeletePolicyRouteRuleSetParamDetail DeletePolicyRouteRuleSet detail param
type DeletePolicyRouteRuleSetParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePolicyRouteRuleSetParam DeletePolicyRouteRuleSet request param
type DeletePolicyRouteRuleSetParam struct {
	BaseParam
	Params DeletePolicyRouteRuleSetParamDetail `json:"deletePolicyRouteRuleSet"`
}
