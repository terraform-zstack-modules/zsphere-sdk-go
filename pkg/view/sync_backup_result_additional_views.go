// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SyncBackupResultView SyncBackupResult
type SyncBackupResultView struct {
	DeletedBackupCount int `json:"deletedBackupCount,omitempty"`
	NewBackupCount int `json:"newBackupCount,omitempty"`
}

