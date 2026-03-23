// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MultipathTopologyStructView MultipathTopologyStruct
type MultipathTopologyStructView struct {
	LunUuid string `json:"lunUuid,omitempty"`
	Devices []DeviceTOView `json:"devices,omitempty"`
}

