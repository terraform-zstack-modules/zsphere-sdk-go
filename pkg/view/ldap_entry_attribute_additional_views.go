// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LdapEntryAttributeInventoryView LdapEntryAttribute
type LdapEntryAttributeInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id string `json:"id,omitempty"`
	Values []interface{} `json:"values,omitempty"`
	OrderMatters bool `json:"orderMatters,omitempty"`
}

