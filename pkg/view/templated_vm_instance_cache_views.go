// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TemplatedVmInstanceCacheInventoryView TemplatedVmInstanceCache
type TemplatedVmInstanceCacheInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	TemplatedVmInstanceUuid string `json:"templatedVmInstanceUuid,omitempty"`
	CacheVmInstanceUuid string `json:"cacheVmInstanceUuid,omitempty"`
}

