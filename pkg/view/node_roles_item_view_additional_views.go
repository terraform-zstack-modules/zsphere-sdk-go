// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NodeRolesItemViewView NodeRolesItemView
type NodeRolesItemViewView struct {
	Uuid string `json:"uuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Role string `json:"role,omitempty"`
}

