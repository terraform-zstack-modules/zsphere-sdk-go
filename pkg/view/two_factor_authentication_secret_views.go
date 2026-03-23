// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TwoFactorAuthenticationSecretInventoryView TwoFactorAuthenticationSecret
type TwoFactorAuthenticationSecretInventoryView struct {
	BaseInfoView
	BaseTimeView
	Secret string `json:"secret,omitempty"`
	Status string `json:"status,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

// GetTwoFactorAuthenticationSecretView GetTwoFactorAuthenticationSecret
type GetTwoFactorAuthenticationSecretView struct {
	Inventory TwoFactorAuthenticationSecretInventoryView `json:"inventory,omitempty"`
}

// QueryTwoFactorAuthenticationView QueryTwoFactorAuthentication
type QueryTwoFactorAuthenticationView struct {
	Inventories []TwoFactorAuthenticationSecretInventoryView `json:"inventories,omitempty"`
}

// ResetTwoFactorAuthenticationSecretEventView ResetTwoFactorAuthenticationSecretEvent
type ResetTwoFactorAuthenticationSecretEventView struct {
	Inventory TwoFactorAuthenticationSecretInventoryView `json:"inventory,omitempty"`
}

