// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NasMountTargetInventoryView NasMountTarget
type NasMountTargetInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	MountDomain string `json:"mountDomain,omitempty"`
	NasFileSystemUuid string `json:"nasFileSystemUuid,omitempty"`
	Type string `json:"type,omitempty"`
}

// UpdateNasMountTargetEventView UpdateNasMountTargetEvent
type UpdateNasMountTargetEventView struct {
	Inventory NasMountTargetInventoryView `json:"inventory,omitempty"`
}

// DeleteNasMountTargetEventView DeleteNasMountTargetEvent
type DeleteNasMountTargetEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryNasMountTargetView QueryNasMountTarget
type QueryNasMountTargetView struct {
	Inventories []NasMountTargetInventoryView `json:"inventories,omitempty"`
}

