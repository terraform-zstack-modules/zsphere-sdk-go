// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeFormatReplyStructView VolumeFormatReplyStruct
type VolumeFormatReplyStructView struct {
	Format string `json:"format,omitempty"`
	MasterHypervisorType string `json:"masterHypervisorType,omitempty"`
	SupportingHypervisorTypes []string `json:"supportingHypervisorTypes,omitempty"`
}

