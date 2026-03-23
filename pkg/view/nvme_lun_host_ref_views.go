// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NvmeLunHostRefInventoryView NvmeLunHostRef
type NvmeLunHostRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	NvmeLunUuid string `json:"nvmeLunUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	Path string `json:"path,omitempty"`
	Hctl string `json:"hctl,omitempty"`
	Locate string `json:"locate,omitempty"`
	Transport string `json:"transport,omitempty"`
}

