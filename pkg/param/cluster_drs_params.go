// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateClusterDRSParamDetail CreateClusterDRS detail param
type CreateClusterDRSParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	AutomationLevel string `json:"automationLevel" validate:"required"`
	Thresholds []ThresholdParam `json:"thresholds" validate:"required"`
	ThresholdDuration int `json:"thresholdDuration" validate:"required"`
	DefaultEnable *bool `json:"defaultEnable,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateClusterDRSParam CreateClusterDRS request param
type CreateClusterDRSParam struct {
	BaseParam
	Params CreateClusterDRSParamDetail `json:"params"`
}
// UpdateClusterDRSParamDetail UpdateClusterDRS detail param
type UpdateClusterDRSParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	AutomationLevel *string `json:"automationLevel,omitempty"`
	Thresholds []ThresholdParam `json:"thresholds,omitempty"`
	ThresholdDuration *int `json:"thresholdDuration,omitempty"`
	State *string `json:"state,omitempty"`
}

// UpdateClusterDRSParam UpdateClusterDRS request param
type UpdateClusterDRSParam struct {
	BaseParam
	Params UpdateClusterDRSParamDetail `json:"updateClusterDRS"`
}
// DeleteClusterDRSParamDetail DeleteClusterDRS detail param
type DeleteClusterDRSParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteClusterDRSParam DeleteClusterDRS request param
type DeleteClusterDRSParam struct {
	BaseParam
	Params DeleteClusterDRSParamDetail `json:"deleteClusterDRS"`
}
