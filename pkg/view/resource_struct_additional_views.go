// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResourceStructView ResourceStruct
type ResourceStructView struct {
	ResourceName string `json:"resourceName,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	DeletePolicy string `json:"deletePolicy,omitempty"`
	Description string `json:"description,omitempty"`
	InDegree []string `json:"inDegree,omitempty"`
	Action string `json:"action,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Results interface{} `json:"results,omitempty"`
	Type string `json:"type,omitempty"`
	Created bool `json:"created,omitempty"`
	MockFailed bool `json:"mockFailed,omitempty"`
}

