// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ScsiLunHostRefInventoryView ScsiLunHostRef
type ScsiLunHostRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ScsiLunUuid string `json:"scsiLunUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	Path string `json:"path,omitempty"`
	Hctl string `json:"hctl,omitempty"`
}

