// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LdapServerInventoryView LdapServer
type LdapServerInventoryView struct {
	BaseInfoView
	BaseTimeView
	Url string `json:"url,omitempty"`
	Base string `json:"base,omitempty"`
	Username string `json:"username,omitempty"`
	ServerType string `json:"serverType,omitempty"`
	Encryption string `json:"encryption,omitempty"`
	Filter string `json:"filter,omitempty"`
	UsernameProperty string `json:"usernameProperty,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	CreateAccountStrategy string `json:"createAccountStrategy,omitempty"`
	DeleteAccountStrategy string `json:"deleteAccountStrategy,omitempty"`
}

// AddLdapServerEventView AddLdapServerEvent
type AddLdapServerEventView struct {
	Inventory LdapServerInventoryView `json:"inventory,omitempty"`
}

// DeleteLdapServerEventView DeleteLdapServerEvent
type DeleteLdapServerEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryLdapServerView QueryLdapServer
type QueryLdapServerView struct {
	Inventories []LdapServerInventoryView `json:"inventories,omitempty"`
}

// UpdateLdapServerEventView UpdateLdapServerEvent
type UpdateLdapServerEventView struct {
	Inventory LdapServerInventoryView `json:"inventory,omitempty"`
}

