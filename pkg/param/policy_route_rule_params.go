// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreatePolicyRouteRuleParamDetail CreatePolicyRouteRule detail param
type CreatePolicyRouteRuleParamDetail struct {
	RuleSetUuid string `json:"ruleSetUuid" validate:"required"`
	TableUuid string `json:"tableUuid" validate:"required"`
	RuleNumber int `json:"ruleNumber" validate:"required"`
	DestIp *string `json:"destIp,omitempty"`
	SourceIp *string `json:"sourceIp,omitempty"`
	DestPort *string `json:"destPort,omitempty"`
	SourcePort *string `json:"sourcePort,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePolicyRouteRuleParam CreatePolicyRouteRule request param
type CreatePolicyRouteRuleParam struct {
	BaseParam
	Params CreatePolicyRouteRuleParamDetail `json:"params"`
}
// DeletePolicyRouteRuleParamDetail DeletePolicyRouteRule detail param
type DeletePolicyRouteRuleParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePolicyRouteRuleParam DeletePolicyRouteRule request param
type DeletePolicyRouteRuleParam struct {
	BaseParam
	Params DeletePolicyRouteRuleParamDetail `json:"deletePolicyRouteRule"`
}
