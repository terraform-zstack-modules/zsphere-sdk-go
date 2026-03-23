// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateSNSWeComEndpointParamDetail UpdateSNSWeComEndpoint detail param
type UpdateSNSWeComEndpointParamDetail struct {
	Url *string `json:"url,omitempty"`
	AtAll *bool `json:"atAll,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
}

// UpdateSNSWeComEndpointParam UpdateSNSWeComEndpoint request param
type UpdateSNSWeComEndpointParam struct {
	BaseParam
	Params UpdateSNSWeComEndpointParamDetail `json:"updateSNSWeComEndpoint"`
}
// CreateSNSWeComEndpointParamDetail CreateSNSWeComEndpoint detail param
type CreateSNSWeComEndpointParamDetail struct {
	Url string `json:"url" validate:"required"`
	AtAll *bool `json:"atAll,omitempty"`
	AtPersonUserIds []string `json:"atPersonUserIds,omitempty"`
	AtPersonList map[string]string `json:"atPersonList,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSWeComEndpointParam CreateSNSWeComEndpoint request param
type CreateSNSWeComEndpointParam struct {
	BaseParam
	Params CreateSNSWeComEndpointParamDetail `json:"params"`
}
