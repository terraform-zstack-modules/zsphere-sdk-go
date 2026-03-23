// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// FailureView Failure
type FailureView struct {
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
}

