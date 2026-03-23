// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateSNSDingTalkEndpointParamDetail UpdateSNSDingTalkEndpoint detail param
type UpdateSNSDingTalkEndpointParamDetail struct {
	Url *string `json:"url,omitempty"`
	AtAll *bool `json:"atAll,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
}

// UpdateSNSDingTalkEndpointParam UpdateSNSDingTalkEndpoint request param
type UpdateSNSDingTalkEndpointParam struct {
	BaseParam
	Params UpdateSNSDingTalkEndpointParamDetail `json:"updateSNSDingTalkEndpoint"`
}
// CreateSNSDingTalkEndpointParamDetail CreateSNSDingTalkEndpoint detail param
type CreateSNSDingTalkEndpointParamDetail struct {
	Url string `json:"url" validate:"required"`
	AtAll *bool `json:"atAll,omitempty"`
	Secret *string `json:"secret,omitempty"`
	AtPersonPhoneNumbers []string `json:"atPersonPhoneNumbers,omitempty"`
	AtPersonList map[string]string `json:"atPersonList,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSDingTalkEndpointParam CreateSNSDingTalkEndpoint request param
type CreateSNSDingTalkEndpointParam struct {
	BaseParam
	Params CreateSNSDingTalkEndpointParamDetail `json:"params"`
}
