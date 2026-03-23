// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateSNSMicrosoftTeamsEndpointParamDetail CreateSNSMicrosoftTeamsEndpoint detail param
type CreateSNSMicrosoftTeamsEndpointParamDetail struct {
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSMicrosoftTeamsEndpointParam CreateSNSMicrosoftTeamsEndpoint request param
type CreateSNSMicrosoftTeamsEndpointParam struct {
	BaseParam
	Params CreateSNSMicrosoftTeamsEndpointParamDetail `json:"params"`
}
// UpdateSNSMicrosoftTeamsEndpointParamDetail UpdateSNSMicrosoftTeamsEndpoint detail param
type UpdateSNSMicrosoftTeamsEndpointParamDetail struct {
	Url *string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
}

// UpdateSNSMicrosoftTeamsEndpointParam UpdateSNSMicrosoftTeamsEndpoint request param
type UpdateSNSMicrosoftTeamsEndpointParam struct {
	BaseParam
	Params UpdateSNSMicrosoftTeamsEndpointParamDetail `json:"updateSNSMicrosoftTeamsEndpoint"`
}
