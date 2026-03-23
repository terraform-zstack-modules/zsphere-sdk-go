// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ChronyServerInfoPairView ChronyServerInfoPair
type ChronyServerInfoPairView struct {
	Internal ChronyServerInfoView `json:"internal,omitempty"`
	External ChronyServerInfoView `json:"external,omitempty"`
}

