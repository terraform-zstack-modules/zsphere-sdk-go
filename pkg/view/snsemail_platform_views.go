// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSEmailPlatformInventoryView SNSEmailPlatform
type SNSEmailPlatformInventoryView struct {
	BaseInfoView
	BaseTimeView
	SmtpServer string `json:"smtpServer,omitempty"`
	SmtpPort int `json:"smtpPort,omitempty"`
	Username string `json:"username,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
}

// ValidateSNSEmailPlatformEventView ValidateSNSEmailPlatformEvent
type ValidateSNSEmailPlatformEventView struct {
	Success bool `json:"success,omitempty"`
}

