// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NodeRolesViewView NodeRolesView
type NodeRolesViewView struct {
	Uuid string `json:"uuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Roles []NodeRolesItemViewView `json:"roles,omitempty"`
}

