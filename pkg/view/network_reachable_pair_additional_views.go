// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NetworkReachablePairView NetworkReachablePair
type NetworkReachablePairView struct {
	SourceHostname string `json:"sourceHostname,omitempty"`
	TargetHostname string `json:"targetHostname,omitempty"`
	Status string `json:"status,omitempty"`
}

