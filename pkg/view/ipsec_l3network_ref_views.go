// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// IPsecL3NetworkRefInventoryView IPsecL3NetworkRef
type IPsecL3NetworkRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ConnectionUuid string `json:"connectionUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
}

