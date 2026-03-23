// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateLoadBalancerParamDetail UpdateLoadBalancer detail param
type UpdateLoadBalancerParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateLoadBalancerParam UpdateLoadBalancer request param
type UpdateLoadBalancerParam struct {
	BaseParam
	Params UpdateLoadBalancerParamDetail `json:"updateLoadBalancer"`
}
// DeleteLoadBalancerParamDetail DeleteLoadBalancer detail param
type DeleteLoadBalancerParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteLoadBalancerParam DeleteLoadBalancer request param
type DeleteLoadBalancerParam struct {
	BaseParam
	Params DeleteLoadBalancerParamDetail `json:"deleteLoadBalancer"`
}
// RefreshLoadBalancerParamDetail RefreshLoadBalancer detail param
type RefreshLoadBalancerParamDetail struct {
}

// RefreshLoadBalancerParam RefreshLoadBalancer request param
type RefreshLoadBalancerParam struct {
	BaseParam
	Params RefreshLoadBalancerParamDetail `json:"refreshLoadBalancer"`
}
// CreateLoadBalancerParamDetail CreateLoadBalancer detail param
type CreateLoadBalancerParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	VipUuid string `json:"vipUuid" validate:"required"`
	Type *string `json:"type,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateLoadBalancerParam CreateLoadBalancer request param
type CreateLoadBalancerParam struct {
	BaseParam
	Params CreateLoadBalancerParamDetail `json:"params"`
}
