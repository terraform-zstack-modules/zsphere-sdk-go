// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// L3NetworkDnsInventoryView L3NetworkDns
type L3NetworkDnsInventoryView struct {
	BaseInfoView
	BaseTimeView
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Dns string `json:"dns,omitempty"`
}

