// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// OvfOSInfoView OvfOSInfo
type OvfOSInfoView struct {
	Id int `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
	OsType string `json:"osType,omitempty"`
	Description string `json:"description,omitempty"`
}

