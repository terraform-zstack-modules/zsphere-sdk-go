// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DetachPortForwardingRuleParamDetail DetachPortForwardingRule detail param
type DetachPortForwardingRuleParamDetail struct {
}

// DetachPortForwardingRuleParam DetachPortForwardingRule request param
type DetachPortForwardingRuleParam struct {
	BaseParam
	Params DetachPortForwardingRuleParamDetail `json:"detachPortForwardingRule"`
}
// DeletePortForwardingRuleParamDetail DeletePortForwardingRule detail param
type DeletePortForwardingRuleParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePortForwardingRuleParam DeletePortForwardingRule request param
type DeletePortForwardingRuleParam struct {
	BaseParam
	Params DeletePortForwardingRuleParamDetail `json:"deletePortForwardingRule"`
}
// AttachPortForwardingRuleParamDetail AttachPortForwardingRule detail param
type AttachPortForwardingRuleParamDetail struct {
}

// AttachPortForwardingRuleParam AttachPortForwardingRule request param
type AttachPortForwardingRuleParam struct {
	BaseParam
	Params AttachPortForwardingRuleParamDetail `json:"params"`
}
// UpdatePortForwardingRuleParamDetail UpdatePortForwardingRule detail param
type UpdatePortForwardingRuleParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdatePortForwardingRuleParam UpdatePortForwardingRule request param
type UpdatePortForwardingRuleParam struct {
	BaseParam
	Params UpdatePortForwardingRuleParamDetail `json:"updatePortForwardingRule"`
}
// CreatePortForwardingRuleParamDetail CreatePortForwardingRule detail param
type CreatePortForwardingRuleParamDetail struct {
	VipUuid string `json:"vipUuid" validate:"required"`
	VipPortStart int `json:"vipPortStart" validate:"required"`
	VipPortEnd *int `json:"vipPortEnd,omitempty"`
	PrivatePortStart *int `json:"privatePortStart,omitempty"`
	PrivatePortEnd *int `json:"privatePortEnd,omitempty"`
	ProtocolType string `json:"protocolType" validate:"required"`
	VmNicUuid *string `json:"vmNicUuid,omitempty"`
	AllowedCidr *string `json:"allowedCidr,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePortForwardingRuleParam CreatePortForwardingRule request param
type CreatePortForwardingRuleParam struct {
	BaseParam
	Params CreatePortForwardingRuleParamDetail `json:"params"`
}
