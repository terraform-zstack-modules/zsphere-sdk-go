// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccountGroupAccountRefInventoryView AccountGroupAccountRef
type AccountGroupAccountRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	GroupUuid string `json:"groupUuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

