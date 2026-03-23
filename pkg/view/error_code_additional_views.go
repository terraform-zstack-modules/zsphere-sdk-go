// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ErrorCodeView ErrorCode
type ErrorCodeView struct {
	Code string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Details string `json:"details,omitempty"`
	I18nDetails string `json:"i18nDetails,omitempty"`
	Elaboration string `json:"elaboration,omitempty"`
	Cause *ErrorCodeView `json:"cause,omitempty"`
	Causes []*ErrorCodeView `json:"causes,omitempty"`
	Opaque map[string]interface{} `json:"opaque,omitempty"`
}

