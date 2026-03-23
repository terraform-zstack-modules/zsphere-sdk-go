// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// OAuth2TokenInventoryView OAuth2Token
type OAuth2TokenInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccessToken string `json:"accessToken,omitempty"`
	IdToken string `json:"idToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	ClientUuid string `json:"clientUuid,omitempty"`
	UserUuid string `json:"userUuid,omitempty"`
}

// GetOAuth2TokenView GetOAuth2Token
type GetOAuth2TokenView struct {
	Inventory OAuth2TokenInventoryView `json:"inventory,omitempty"`
}

