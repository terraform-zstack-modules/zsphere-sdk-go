// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// ValidateSNSEmailPlatformParamDetail ValidateSNSEmailPlatform detail param
type ValidateSNSEmailPlatformParamDetail struct {
	SmtpServer *string `json:"smtpServer,omitempty"`
	SmtpPort *int `json:"smtpPort,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

// ValidateSNSEmailPlatformParam ValidateSNSEmailPlatform request param
type ValidateSNSEmailPlatformParam struct {
	BaseParam
	Params ValidateSNSEmailPlatformParamDetail `json:"validateSNSEmailPlatform"`
}
// CreateSNSEmailPlatformParamDetail CreateSNSEmailPlatform detail param
type CreateSNSEmailPlatformParamDetail struct {
	SmtpServer string `json:"smtpServer" validate:"required"`
	SmtpPort int `json:"smtpPort" validate:"required"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	EncryptType *string `json:"encryptType,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSEmailPlatformParam CreateSNSEmailPlatform request param
type CreateSNSEmailPlatformParam struct {
	BaseParam
	Params CreateSNSEmailPlatformParamDetail `json:"params"`
}
