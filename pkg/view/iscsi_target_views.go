// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// IscsiTargetInventoryView IscsiTarget
type IscsiTargetInventoryView struct {
	BaseInfoView
	BaseTimeView
	IscsiServerUuid string `json:"iscsiServerUuid,omitempty"`
	Iqn string `json:"iqn,omitempty"`
	IscsiLuns []IscsiLunInventoryView `json:"iscsiLuns,omitempty"`
}

