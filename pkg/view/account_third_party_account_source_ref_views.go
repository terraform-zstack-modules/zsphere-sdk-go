// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccountThirdPartyAccountSourceRefInventoryView AccountThirdPartyAccountSourceRef
type AccountThirdPartyAccountSourceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	Credentials string `json:"credentials,omitempty"`
	AccountSourceUuid string `json:"accountSourceUuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

// QueryThirdPartyAccountSourceBindingView QueryThirdPartyAccountSourceBinding
type QueryThirdPartyAccountSourceBindingView struct {
	Inventories []AccountThirdPartyAccountSourceRefInventoryView `json:"inventories,omitempty"`
}

// CreateLdapBindingEventView CreateLdapBindingEvent
type CreateLdapBindingEventView struct {
	Inventory AccountThirdPartyAccountSourceRefInventoryView `json:"inventory,omitempty"`
}

