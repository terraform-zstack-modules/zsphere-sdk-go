// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateSSORedirectTemplateParamDetail CreateSSORedirectTemplate detail param
type CreateSSORedirectTemplateParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description" validate:"required"`
	ClientUuid string `json:"clientUuid" validate:"required"`
	RedirectTemplate string `json:"redirectTemplate" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSSORedirectTemplateParam CreateSSORedirectTemplate request param
type CreateSSORedirectTemplateParam struct {
	BaseParam
	Params CreateSSORedirectTemplateParamDetail `json:"params"`
}
// UpdateSSORedirectTemplateParamDetail UpdateSSORedirectTemplate detail param
type UpdateSSORedirectTemplateParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	RedirectTemplate string `json:"redirectTemplate" validate:"required"`
}

// UpdateSSORedirectTemplateParam UpdateSSORedirectTemplate request param
type UpdateSSORedirectTemplateParam struct {
	BaseParam
	Params UpdateSSORedirectTemplateParamDetail `json:"params"`
}
// DeleteSSORedirectTemplateParamDetail DeleteSSORedirectTemplate detail param
type DeleteSSORedirectTemplateParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteSSORedirectTemplateParam DeleteSSORedirectTemplate request param
type DeleteSSORedirectTemplateParam struct {
	BaseParam
	Params DeleteSSORedirectTemplateParamDetail `json:"params"`
}
