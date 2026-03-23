// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteHostKernelInterfaceParamDetail DeleteHostKernelInterface detail param
type DeleteHostKernelInterfaceParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteHostKernelInterfaceParam DeleteHostKernelInterface request param
type DeleteHostKernelInterfaceParam struct {
	BaseParam
	Params DeleteHostKernelInterfaceParamDetail `json:"deleteHostKernelInterface"`
}
// CreateHostKernelInterfaceParamDetail CreateHostKernelInterface detail param
type CreateHostKernelInterfaceParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	HostUuid string `json:"hostUuid" validate:"required"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	RequiredIp *string `json:"requiredIp,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	TrafficTypes []string `json:"trafficTypes,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateHostKernelInterfaceParam CreateHostKernelInterface request param
type CreateHostKernelInterfaceParam struct {
	BaseParam
	Params CreateHostKernelInterfaceParamDetail `json:"params"`
}
// UpdateHostKernelInterfaceParamDetail UpdateHostKernelInterface detail param
type UpdateHostKernelInterfaceParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	RequiredIp *string `json:"requiredIp,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	TrafficTypes []string `json:"trafficTypes,omitempty"`
}

// UpdateHostKernelInterfaceParam UpdateHostKernelInterface request param
type UpdateHostKernelInterfaceParam struct {
	BaseParam
	Params UpdateHostKernelInterfaceParamDetail `json:"updateHostKernelInterface"`
}
