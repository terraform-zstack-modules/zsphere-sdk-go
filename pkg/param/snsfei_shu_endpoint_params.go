// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateSNSFeiShuEndpointParamDetail UpdateSNSFeiShuEndpoint detail param
type UpdateSNSFeiShuEndpointParamDetail struct {
	Url *string `json:"url,omitempty"`
	AtAll *bool `json:"atAll,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
}

// UpdateSNSFeiShuEndpointParam UpdateSNSFeiShuEndpoint request param
type UpdateSNSFeiShuEndpointParam struct {
	BaseParam
	Params UpdateSNSFeiShuEndpointParamDetail `json:"updateSNSFeiShuEndpoint"`
}
// CreateSNSFeiShuEndpointParamDetail CreateSNSFeiShuEndpoint detail param
type CreateSNSFeiShuEndpointParamDetail struct {
	Url string `json:"url" validate:"required"`
	AtAll *bool `json:"atAll,omitempty"`
	AtPersonUserIds []string `json:"atPersonUserIds,omitempty"`
	Secret *string `json:"secret,omitempty"`
	AtPersonList map[string]string `json:"atPersonList,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSFeiShuEndpointParam CreateSNSFeiShuEndpoint request param
type CreateSNSFeiShuEndpointParam struct {
	BaseParam
	Params CreateSNSFeiShuEndpointParamDetail `json:"params"`
}
