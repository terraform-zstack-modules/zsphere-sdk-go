// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteFlowCollectorParamDetail DeleteFlowCollector detail param
type DeleteFlowCollectorParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteFlowCollectorParam DeleteFlowCollector request param
type DeleteFlowCollectorParam struct {
	BaseParam
	Params DeleteFlowCollectorParamDetail `json:"deleteFlowCollector"`
}
// CreateFlowCollectorParamDetail CreateFlowCollector detail param
type CreateFlowCollectorParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Server *string `json:"server,omitempty"`
	Port *int64 `json:"port,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFlowCollectorParam CreateFlowCollector request param
type CreateFlowCollectorParam struct {
	BaseParam
	Params CreateFlowCollectorParamDetail `json:"params"`
}
// UpdateFlowCollectorParamDetail UpdateFlowCollector detail param
type UpdateFlowCollectorParamDetail struct {
	Server *string `json:"server,omitempty"`
	Port *int64 `json:"port,omitempty"`
}

// UpdateFlowCollectorParam UpdateFlowCollector request param
type UpdateFlowCollectorParam struct {
	BaseParam
	Params UpdateFlowCollectorParamDetail `json:"updateFlowCollector"`
}
