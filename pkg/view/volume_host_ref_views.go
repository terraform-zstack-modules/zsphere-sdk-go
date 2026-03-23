// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeHostRefInventoryView VolumeHostRef
type VolumeHostRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostUuid string `json:"hostUuid,omitempty"`
	VolumeUuid string `json:"volumeUuid,omitempty"`
	MountPath string `json:"mountPath,omitempty"`
	Device string `json:"device,omitempty"`
}

