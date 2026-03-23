// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ChronyServerInfoView ChronyServerInfo
type ChronyServerInfoView struct {
	Hostname string `json:"hostname,omitempty"`
	Status string `json:"status,omitempty"`
}

