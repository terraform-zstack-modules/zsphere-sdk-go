// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AddHostFromFileResultView AddHostFromFileResult
type AddHostFromFileResultView struct {
	Ip string `json:"ip,omitempty"`
	Success bool `json:"success,omitempty"`
}

