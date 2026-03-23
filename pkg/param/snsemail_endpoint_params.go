// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateSNSEmailEndpointParamDetail CreateSNSEmailEndpoint detail param
type CreateSNSEmailEndpointParamDetail struct {
	Email *string `json:"email,omitempty"`
	Emails []string `json:"emails,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSEmailEndpointParam CreateSNSEmailEndpoint request param
type CreateSNSEmailEndpointParam struct {
	BaseParam
	Params CreateSNSEmailEndpointParamDetail `json:"params"`
}
