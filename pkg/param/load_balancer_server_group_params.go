// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteLoadBalancerServerGroupParamDetail DeleteLoadBalancerServerGroup detail param
type DeleteLoadBalancerServerGroupParamDetail struct {
}

// DeleteLoadBalancerServerGroupParam DeleteLoadBalancerServerGroup request param
type DeleteLoadBalancerServerGroupParam struct {
	BaseParam
	Params DeleteLoadBalancerServerGroupParamDetail `json:"deleteLoadBalancerServerGroup"`
}
// CreateLoadBalancerServerGroupParamDetail CreateLoadBalancerServerGroup detail param
type CreateLoadBalancerServerGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateLoadBalancerServerGroupParam CreateLoadBalancerServerGroup request param
type CreateLoadBalancerServerGroupParam struct {
	BaseParam
	Params CreateLoadBalancerServerGroupParamDetail `json:"params"`
}
// UpdateLoadBalancerServerGroupParamDetail UpdateLoadBalancerServerGroup detail param
type UpdateLoadBalancerServerGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateLoadBalancerServerGroupParam UpdateLoadBalancerServerGroup request param
type UpdateLoadBalancerServerGroupParam struct {
	BaseParam
	Params UpdateLoadBalancerServerGroupParamDetail `json:"updateLoadBalancerServerGroup"`
}
