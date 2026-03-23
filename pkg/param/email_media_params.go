// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateEmailMediaParamDetail CreateEmailMedia detail param
type CreateEmailMediaParamDetail struct {
	SmtpServer string `json:"smtpServer" validate:"required"`
	SmtpPort int `json:"smtpPort" validate:"required"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEmailMediaParam CreateEmailMedia request param
type CreateEmailMediaParam struct {
	BaseParam
	Params CreateEmailMediaParamDetail `json:"params"`
}
// UpdateEmailMediaParamDetail UpdateEmailMedia detail param
type UpdateEmailMediaParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	SmtpServer *string `json:"smtpServer,omitempty"`
	SmtpPort *int `json:"smtpPort,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

// UpdateEmailMediaParam UpdateEmailMedia request param
type UpdateEmailMediaParam struct {
	BaseParam
	Params UpdateEmailMediaParamDetail `json:"updateEmailMedia"`
}
