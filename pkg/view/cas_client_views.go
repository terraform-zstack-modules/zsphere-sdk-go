// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CasClientInventoryView CasClient
type CasClientInventoryView struct {
	BaseInfoView
	BaseTimeView
	LoginMNUrl string `json:"loginMNUrl,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
	CasServerLoginUrl string `json:"casServerLoginUrl,omitempty"`
	CasServerUrlPrefix string `json:"casServerUrlPrefix,omitempty"`
	ServerName string `json:"serverName,omitempty"`
	State string `json:"state,omitempty"`
	UsernameProperty string `json:"usernameProperty,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	CreateAccountStrategy string `json:"createAccountStrategy,omitempty"`
	DeleteAccountStrategy string `json:"deleteAccountStrategy,omitempty"`
}

// CreateCasClientEventView CreateCasClientEvent
type CreateCasClientEventView struct {
	Inventory CasClientInventoryView `json:"inventory,omitempty"`
}

// UpdateCasClientEventView UpdateCasClientEvent
type UpdateCasClientEventView struct {
	Inventory CasClientInventoryView `json:"inventory,omitempty"`
}

