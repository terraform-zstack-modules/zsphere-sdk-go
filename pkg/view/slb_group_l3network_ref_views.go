// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SlbGroupL3NetworkRefInventoryView SlbGroupL3NetworkRef
type SlbGroupL3NetworkRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	SlbGroupUuid string `json:"slbGroupUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	L3NetworkCategory string `json:"l3NetworkCategory,omitempty"`
	L3NetworkType string `json:"l3NetworkType,omitempty"`
	Type string `json:"type,omitempty"`
}

