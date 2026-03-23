// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PowerOffHardwareResultView PowerOffHardwareResult
type PowerOffHardwareResultView struct {
	Uuid string `json:"uuid,omitempty"`
	Success bool `json:"success,omitempty"`
}

