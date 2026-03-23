// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// ChangeVmNicSecurityPolicyParamDetail ChangeVmNicSecurityPolicy detail param
type ChangeVmNicSecurityPolicyParamDetail struct {
	IngressPolicy *string `json:"ingressPolicy,omitempty"`
	EgressPolicy *string `json:"egressPolicy,omitempty"`
}

// ChangeVmNicSecurityPolicyParam ChangeVmNicSecurityPolicy request param
type ChangeVmNicSecurityPolicyParam struct {
	BaseParam
	Params ChangeVmNicSecurityPolicyParamDetail `json:"changeVmNicSecurityPolicy"`
}
