// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateSNSHttpEndpointParamDetail CreateSNSHttpEndpoint detail param
type CreateSNSHttpEndpointParamDetail struct {
	Url string `json:"url" validate:"required"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSHttpEndpointParam CreateSNSHttpEndpoint request param
type CreateSNSHttpEndpointParam struct {
	BaseParam
	Params CreateSNSHttpEndpointParamDetail `json:"params"`
}
// UpdateSNSHttpEndpointParamDetail UpdateSNSHttpEndpoint detail param
type UpdateSNSHttpEndpointParamDetail struct {
	Url *string `json:"url,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
}

// UpdateSNSHttpEndpointParam UpdateSNSHttpEndpoint request param
type UpdateSNSHttpEndpointParam struct {
	BaseParam
	Params UpdateSNSHttpEndpointParamDetail `json:"updateSNSHttpEndpoint"`
}
