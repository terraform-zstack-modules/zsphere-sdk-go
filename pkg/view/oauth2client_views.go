// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// OAuth2ClientInventoryView OAuth2Client
type OAuth2ClientInventoryView struct {
	BaseInfoView
	BaseTimeView
	ClientId string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	GrantType string `json:"grantType,omitempty"`
	LoginMNUrl string `json:"loginMNUrl,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
	AuthorizationUrl string `json:"authorizationUrl,omitempty"`
	TokenUrl string `json:"tokenUrl,omitempty"`
	UserinfoUrl string `json:"userinfoUrl,omitempty"`
	LogoutUrl string `json:"logoutUrl,omitempty"`
	UsernameProperty string `json:"usernameProperty,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	CreateAccountStrategy string `json:"createAccountStrategy,omitempty"`
	DeleteAccountStrategy string `json:"deleteAccountStrategy,omitempty"`
}

// CreateOAuthClientEventView CreateOAuthClientEvent
type CreateOAuthClientEventView struct {
	Inventory OAuth2ClientInventoryView `json:"inventory,omitempty"`
}

// UpdateOAuthClientEventView UpdateOAuthClientEvent
type UpdateOAuthClientEventView struct {
	Inventory OAuth2ClientInventoryView `json:"inventory,omitempty"`
}

