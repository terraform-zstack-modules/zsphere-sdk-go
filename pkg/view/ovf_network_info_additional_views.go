// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// OvfNetworkInfoView OvfNetworkInfo
type OvfNetworkInfoView struct {
	Name string `json:"name,omitempty"`
}

