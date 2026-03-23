// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SSOTokenInventoryView SSOToken
type SSOTokenInventoryView struct {
	BaseInfoView
	BaseTimeView
	ClientUuid string `json:"clientUuid,omitempty"`
	UserUuid string `json:"userUuid,omitempty"`
}

