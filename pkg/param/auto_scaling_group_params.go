// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateAutoScalingGroupParamDetail UpdateAutoScalingGroup detail param
type UpdateAutoScalingGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	MinResourceSize *int `json:"minResourceSize,omitempty"`
	MaxResourceSize *int `json:"maxResourceSize,omitempty"`
	RemovalPolicy *string `json:"removalPolicy,omitempty"`
}

// UpdateAutoScalingGroupParam UpdateAutoScalingGroup request param
type UpdateAutoScalingGroupParam struct {
	BaseParam
	Params UpdateAutoScalingGroupParamDetail `json:"updateAutoScalingGroup"`
}
// DeleteAutoScalingGroupParamDetail DeleteAutoScalingGroup detail param
type DeleteAutoScalingGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingGroupParam DeleteAutoScalingGroup request param
type DeleteAutoScalingGroupParam struct {
	BaseParam
	Params DeleteAutoScalingGroupParamDetail `json:"deleteAutoScalingGroup"`
}
// CreateAutoScalingGroupParamDetail CreateAutoScalingGroup detail param
type CreateAutoScalingGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ScalingResourceType string `json:"scalingResourceType" validate:"required"`
	MinResourceSize int `json:"minResourceSize" validate:"required"`
	MaxResourceSize int `json:"maxResourceSize" validate:"required"`
	DefaultCooldown int64 `json:"defaultCooldown" validate:"required"`
	RemovalPolicy string `json:"removalPolicy" validate:"required"`
	DefaultEnable *bool `json:"defaultEnable,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingGroupParam CreateAutoScalingGroup request param
type CreateAutoScalingGroupParam struct {
	BaseParam
	Params CreateAutoScalingGroupParamDetail `json:"params"`
}
