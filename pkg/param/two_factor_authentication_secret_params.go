// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// GetTwoFactorAuthenticationSecretParamDetail GetTwoFactorAuthenticationSecret detail param
type GetTwoFactorAuthenticationSecretParamDetail struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	CaptchaUuid *string `json:"captchaUuid,omitempty"`
	VerifyCode *string `json:"verifyCode,omitempty"`
}

// GetTwoFactorAuthenticationSecretParam GetTwoFactorAuthenticationSecret request param
type GetTwoFactorAuthenticationSecretParam struct {
	BaseParam
	Params GetTwoFactorAuthenticationSecretParamDetail `json:"getTwoFactorAuthenticationSecret"`
}
// ResetTwoFactorAuthenticationSecretParamDetail ResetTwoFactorAuthenticationSecret detail param
type ResetTwoFactorAuthenticationSecretParamDetail struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	CaptchaUuid *string `json:"captchaUuid,omitempty"`
	VerifyCode *string `json:"verifyCode,omitempty"`
}

// ResetTwoFactorAuthenticationSecretParam ResetTwoFactorAuthenticationSecret request param
type ResetTwoFactorAuthenticationSecretParam struct {
	BaseParam
	Params ResetTwoFactorAuthenticationSecretParamDetail `json:"resetTwoFactorAuthenticationSecret"`
}
