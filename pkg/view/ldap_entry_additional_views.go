// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LdapEntryInventoryView LdapEntry
type LdapEntryInventoryView struct {
	BaseInfoView
	BaseTimeView
	Dn string `json:"dn,omitempty"`
	Enable bool `json:"enable,omitempty"`
	Attributes []LdapEntryAttributeInventoryView `json:"attributes,omitempty"`
}

