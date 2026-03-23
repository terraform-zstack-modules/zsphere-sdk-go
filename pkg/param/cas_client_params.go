// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateCasClientParamDetail CreateCasClient detail param
type CreateCasClientParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	CasServerLoginUrl string `json:"casServerLoginUrl" validate:"required"`
	CasServerUrlPrefix string `json:"casServerUrlPrefix" validate:"required"`
	ServerName string `json:"serverName" validate:"required"`
	UsernameProperty *string `json:"usernameProperty,omitempty"`
	UrlTemplate *string `json:"urlTemplate,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateCasClientParam CreateCasClient request param
type CreateCasClientParam struct {
	BaseParam
	Params CreateCasClientParamDetail `json:"params"`
}
// UpdateCasClientParamDetail UpdateCasClient detail param
type UpdateCasClientParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Description *string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	CasServerLoginUrl *string `json:"casServerLoginUrl,omitempty"`
	CasServerUrlPrefix *string `json:"casServerUrlPrefix,omitempty"`
	ServerName *string `json:"serverName,omitempty"`
	UsernameProperty *string `json:"usernameProperty,omitempty"`
}

// UpdateCasClientParam UpdateCasClient request param
type UpdateCasClientParam struct {
	BaseParam
	Params UpdateCasClientParamDetail `json:"params"`
}
