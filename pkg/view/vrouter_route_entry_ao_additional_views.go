// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VRouterRouteEntryAOView VRouterRouteEntryAO
type VRouterRouteEntryAOView struct {
	Destination string `json:"destination,omitempty"`
	Target string `json:"target,omitempty"`
	Type string `json:"type,omitempty"`
	Status string `json:"status,omitempty"`
	Distance int `json:"distance,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
}

