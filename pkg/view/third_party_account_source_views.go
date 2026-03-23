// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ThirdPartyAccountSourceInventoryView ThirdPartyAccountSource
type ThirdPartyAccountSourceInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	CreateAccountStrategy string `json:"createAccountStrategy,omitempty"`
	DeleteAccountStrategy string `json:"deleteAccountStrategy,omitempty"`
}

// GetSSOClientView GetSSOClient
type GetSSOClientView struct {
	Inventories []ThirdPartyAccountSourceInventoryView `json:"inventories,omitempty"`
}

