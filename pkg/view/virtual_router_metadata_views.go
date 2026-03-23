// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VirtualRouterMetadataInventoryView VirtualRouterMetadata
type VirtualRouterMetadataInventoryView struct {
	BaseInfoView
	BaseTimeView
	ZvrVersion string `json:"zvrVersion,omitempty"`
	VyosVersion string `json:"vyosVersion,omitempty"`
	KernelVersion string `json:"kernelVersion,omitempty"`
}

