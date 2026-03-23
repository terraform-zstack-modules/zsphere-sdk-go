// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmNicConflictEntryView VmNicConflictEntry
type VmNicConflictEntryView struct {
	Ip string `json:"ip,omitempty"`
	Mac string `json:"mac,omitempty"`
	VmNicName string `json:"vmNicName,omitempty"`
	VmInstanceName string `json:"vmInstanceName,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
}

