// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TemplatedVmInstanceRefInventoryView TemplatedVmInstanceRef
type TemplatedVmInstanceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	TemplatedVmInstanceUuid string `json:"templatedVmInstanceUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
}

