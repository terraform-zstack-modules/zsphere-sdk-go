// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// OvfMemoryInfoView OvfMemoryInfo
type OvfMemoryInfoView struct {
	InstanceId string `json:"instanceId,omitempty"`
	Quantity int64 `json:"quantity,omitempty"`
}

