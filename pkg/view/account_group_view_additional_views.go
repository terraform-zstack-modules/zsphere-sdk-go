// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccountGroupViewView AccountGroupView
type AccountGroupViewView struct {
	GroupUuid string `json:"groupUuid,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	Inventory AccountGroupInventoryView `json:"inventory,omitempty"`
	Accounts []AccountInventoryView `json:"accounts,omitempty"`
	Groups []*AccountGroupViewView `json:"groups,omitempty"`
}

