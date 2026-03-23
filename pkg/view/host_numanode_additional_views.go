// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostNUMANodeView HostNUMANode
type HostNUMANodeView struct {
	Distance []string `json:"distance,omitempty"`
	Cpus []string `json:"cpus,omitempty"`
	Free int64 `json:"free,omitempty"`
	Size int64 `json:"size,omitempty"`
	NodeID string `json:"nodeID,omitempty"`
	VMsUuid []string `json:"VMsUuid,omitempty"`
}

