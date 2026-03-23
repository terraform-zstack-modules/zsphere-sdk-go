// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ManagementsStatusViewView ManagementsStatusView
type ManagementsStatusViewView struct {
	Vip string `json:"vip,omitempty"`
	UiHttpPath string `json:"uiHttpPath,omitempty"`
	Nodes []ManagementNodeStatusViewView `json:"nodes,omitempty"`
}

