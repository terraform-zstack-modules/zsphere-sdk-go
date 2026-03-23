// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateLoadBalancerListenerParamDetail UpdateLoadBalancerListener detail param
type UpdateLoadBalancerListenerParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateLoadBalancerListenerParam UpdateLoadBalancerListener request param
type UpdateLoadBalancerListenerParam struct {
	BaseParam
	Params UpdateLoadBalancerListenerParamDetail `json:"updateLoadBalancerListener"`
}
// ChangeLoadBalancerListenerParamDetail ChangeLoadBalancerListener detail param
type ChangeLoadBalancerListenerParamDetail struct {
	ConnectionIdleTimeout *int `json:"connectionIdleTimeout,omitempty"`
	MaxConnection *int `json:"maxConnection,omitempty"`
	BalancerAlgorithm *string `json:"balancerAlgorithm,omitempty"`
	HealthCheckTarget *string `json:"healthCheckTarget,omitempty"`
	HealthyThreshold *int `json:"healthyThreshold,omitempty"`
	UnhealthyThreshold *int `json:"unhealthyThreshold,omitempty"`
	HealthCheckInterval *int `json:"healthCheckInterval,omitempty"`
	HealthCheckProtocol *string `json:"healthCheckProtocol,omitempty"`
	HealthCheckMethod *string `json:"healthCheckMethod,omitempty"`
	HealthCheckURI *string `json:"healthCheckURI,omitempty"`
	HealthCheckHttpCode *string `json:"healthCheckHttpCode,omitempty"`
	AclStatus *string `json:"aclStatus,omitempty"`
	SecurityPolicyType *string `json:"securityPolicyType,omitempty"`
	Nbprocess *int `json:"nbprocess,omitempty"`
	HttpMode *string `json:"httpMode,omitempty"`
	SessionPersistence *string `json:"sessionPersistence,omitempty"`
	SessionIdleTimeout *int `json:"sessionIdleTimeout,omitempty"`
	CookieName *string `json:"cookieName,omitempty"`
	HttpRedirectHttps *string `json:"httpRedirectHttps,omitempty"`
	RedirectPort *int `json:"redirectPort,omitempty"`
	StatusCode *int `json:"statusCode,omitempty"`
}

// ChangeLoadBalancerListenerParam ChangeLoadBalancerListener request param
type ChangeLoadBalancerListenerParam struct {
	BaseParam
	Params ChangeLoadBalancerListenerParamDetail `json:"changeLoadBalancerListener"`
}
// CreateLoadBalancerListenerParamDetail CreateLoadBalancerListener detail param
type CreateLoadBalancerListenerParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	InstancePort *int `json:"instancePort,omitempty"`
	LoadBalancerPort int `json:"loadBalancerPort" validate:"required"`
	Protocol *string `json:"protocol,omitempty"`
	CertificateUuid *string `json:"certificateUuid,omitempty"`
	HealthCheckProtocol *string `json:"healthCheckProtocol,omitempty"`
	HealthCheckMethod *string `json:"healthCheckMethod,omitempty"`
	HealthCheckURI *string `json:"healthCheckURI,omitempty"`
	HealthCheckHttpCode *string `json:"healthCheckHttpCode,omitempty"`
	AclStatus *string `json:"aclStatus,omitempty"`
	AclUuids []string `json:"aclUuids,omitempty"`
	AclType *string `json:"aclType,omitempty"`
	SecurityPolicyType *string `json:"securityPolicyType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateLoadBalancerListenerParam CreateLoadBalancerListener request param
type CreateLoadBalancerListenerParam struct {
	BaseParam
	Params CreateLoadBalancerListenerParamDetail `json:"params"`
}
// DeleteLoadBalancerListenerParamDetail DeleteLoadBalancerListener detail param
type DeleteLoadBalancerListenerParamDetail struct {
}

// DeleteLoadBalancerListenerParam DeleteLoadBalancerListener request param
type DeleteLoadBalancerListenerParam struct {
	BaseParam
	Params DeleteLoadBalancerListenerParamDetail `json:"deleteLoadBalancerListener"`
}
