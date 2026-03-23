// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// ChangeSecurityGroupRuleParamDetail ChangeSecurityGroupRule detail param
type ChangeSecurityGroupRuleParamDetail struct {
	Description *string `json:"description,omitempty"`
	RemoteSecurityGroupUuid *string `json:"remoteSecurityGroupUuid,omitempty"`
	Action *string `json:"action,omitempty"`
	State *string `json:"state,omitempty"`
	Priority *int `json:"priority,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	SrcIpRange *string `json:"srcIpRange,omitempty"`
	DstIpRange *string `json:"dstIpRange,omitempty"`
	DstPortRange *string `json:"dstPortRange,omitempty"`
}

// ChangeSecurityGroupRuleParam ChangeSecurityGroupRule request param
type ChangeSecurityGroupRuleParam struct {
	BaseParam
	Params ChangeSecurityGroupRuleParamDetail `json:"changeSecurityGroupRule"`
}
// ValidateSecurityGroupRuleParamDetail ValidateSecurityGroupRule detail param
type ValidateSecurityGroupRuleParamDetail struct {
	Type string `json:"type" validate:"required"`
	Protocol string `json:"protocol" validate:"required"`
	RemoteSecurityGroupUuid *string `json:"remoteSecurityGroupUuid,omitempty"`
	IpVersion *int `json:"ipVersion,omitempty"`
	SrcIpRange *string `json:"srcIpRange,omitempty"`
	DstIpRange *string `json:"dstIpRange,omitempty"`
	DstPortRange *string `json:"dstPortRange,omitempty"`
	Action *string `json:"action,omitempty"`
	StartPort *int `json:"startPort,omitempty"`
	EndPort *int `json:"endPort,omitempty"`
	AllowedCidr *string `json:"allowedCidr,omitempty"`
}

// ValidateSecurityGroupRuleParam ValidateSecurityGroupRule request param
type ValidateSecurityGroupRuleParam struct {
	BaseParam
	Params ValidateSecurityGroupRuleParamDetail `json:"validateSecurityGroupRule"`
}
// AddSecurityGroupRuleParamDetail AddSecurityGroupRule detail param
type AddSecurityGroupRuleParamDetail struct {
	Rules []AddSecurityGroupRule_SecurityGroupRuleAOParam `json:"rules" validate:"required"`
	RemoteSecurityGroupUuids []string `json:"remoteSecurityGroupUuids,omitempty"`
	Priority *int `json:"priority,omitempty"`
}

// AddSecurityGroupRuleParam AddSecurityGroupRule request param
type AddSecurityGroupRuleParam struct {
	BaseParam
	Params AddSecurityGroupRuleParamDetail `json:"params"`
}
// DeleteSecurityGroupRuleParamDetail DeleteSecurityGroupRule detail param
type DeleteSecurityGroupRuleParamDetail struct {
	RuleUuids []string `json:"ruleUuids" validate:"required"`
}

// DeleteSecurityGroupRuleParam DeleteSecurityGroupRule request param
type DeleteSecurityGroupRuleParam struct {
	BaseParam
	Params DeleteSecurityGroupRuleParamDetail `json:"deleteSecurityGroupRule"`
}
