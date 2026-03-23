// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteSNSApplicationEndpointParamDetail DeleteSNSApplicationEndpoint detail param
type DeleteSNSApplicationEndpointParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteSNSApplicationEndpointParam DeleteSNSApplicationEndpoint request param
type DeleteSNSApplicationEndpointParam struct {
	BaseParam
	Params DeleteSNSApplicationEndpointParamDetail `json:"deleteSNSApplicationEndpoint"`
}
// UpdateSNSApplicationEndpointParamDetail UpdateSNSApplicationEndpoint detail param
type UpdateSNSApplicationEndpointParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
}

// UpdateSNSApplicationEndpointParam UpdateSNSApplicationEndpoint request param
type UpdateSNSApplicationEndpointParam struct {
	BaseParam
	Params UpdateSNSApplicationEndpointParamDetail `json:"updateSNSApplicationEndpoint"`
}
