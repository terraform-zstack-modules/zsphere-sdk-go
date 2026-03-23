// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// IPsecPeerCidrInventoryView IPsecPeerCidr
type IPsecPeerCidrInventoryView struct {
	BaseInfoView
	BaseTimeView
	Cidr string `json:"cidr,omitempty"`
	ConnectionUuid string `json:"connectionUuid,omitempty"`
}

