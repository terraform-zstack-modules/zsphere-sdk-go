// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SharedBlockCandidateStructView SharedBlockCandidateStruct
type SharedBlockCandidateStructView struct {
	Wwid string `json:"wwid,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Model string `json:"model,omitempty"`
	Wwn string `json:"wwn,omitempty"`
	Serial string `json:"serial,omitempty"`
	Hctl string `json:"hctl,omitempty"`
	Type string `json:"type,omitempty"`
	Path string `json:"path,omitempty"`
	Size int64 `json:"size,omitempty"`
	Source string `json:"source,omitempty"`
	Transport string `json:"transport,omitempty"`
	TargetIdentifier string `json:"targetIdentifier,omitempty"`
}

