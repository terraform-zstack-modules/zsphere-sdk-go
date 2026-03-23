// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccountGroupResourceRefInventoryView AccountGroupResourceRef
type AccountGroupResourceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	GroupUuid string `json:"groupUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
}

