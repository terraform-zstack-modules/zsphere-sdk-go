// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EmailMediaInventoryView EmailMedia
type EmailMediaInventoryView struct {
	BaseInfoView
	BaseTimeView
	SmtpServer string `json:"smtpServer,omitempty"`
	SmtpPort int `json:"smtpPort,omitempty"`
	Username string `json:"username,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateEmailMediaEventView UpdateEmailMediaEvent
type UpdateEmailMediaEventView struct {
	Inventory EmailMediaInventoryView `json:"inventory,omitempty"`
}

