// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccountGroupResourceViewView AccountGroupResourceView
type AccountGroupResourceViewView struct {
	GroupUuid string `json:"groupUuid,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	Resources []ResourceInventoryView `json:"resources,omitempty"`
}

