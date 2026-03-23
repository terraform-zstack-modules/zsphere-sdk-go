// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccessControlListEntryInventoryView AccessControlListEntry
type AccessControlListEntryInventoryView struct {
	BaseInfoView
	BaseTimeView
	AclUuid string `json:"aclUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Domain string `json:"domain,omitempty"`
	Url string `json:"url,omitempty"`
	IpEntries string `json:"ipEntries,omitempty"`
	Description string `json:"description,omitempty"`
}

// RemoveAccessControlListEntryEventView RemoveAccessControlListEntryEvent
type RemoveAccessControlListEntryEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddAccessControlListEntryEventView AddAccessControlListEntryEvent
type AddAccessControlListEntryEventView struct {
	Inventory AccessControlListEntryInventoryView `json:"inventory,omitempty"`
}

// ChangeAccessControlListRedirectRuleEventView ChangeAccessControlListRedirectRuleEvent
type ChangeAccessControlListRedirectRuleEventView struct {
	Inventory AccessControlListEntryInventoryView `json:"inventory,omitempty"`
}

