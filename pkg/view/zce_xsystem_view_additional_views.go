// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZceXSystemViewView ZceXSystemView
type ZceXSystemViewView struct {
	AdminUserName string `json:"adminUserName,omitempty"`
	AdminUserId string `json:"adminUserId,omitempty"`
	Version string `json:"version,omitempty"`
}

