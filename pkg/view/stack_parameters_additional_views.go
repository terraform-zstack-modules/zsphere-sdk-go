// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// StackParametersView StackParameters
type StackParametersView struct {
	ParamName string `json:"paramName,omitempty"`
	Type string `json:"type,omitempty"`
	DefaultValue string `json:"defaultValue,omitempty"`
	Description string `json:"description,omitempty"`
	NoEcho bool `json:"noEcho,omitempty"`
	Label string `json:"label,omitempty"`
	ConstraintDescription string `json:"constraintDescription,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}

