// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EventStructView EventStruct
type EventStructView struct {
	Namespace string `json:"namespace,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	LabelNames []string `json:"labelNames,omitempty"`
}

