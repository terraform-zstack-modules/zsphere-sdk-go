// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteResourceAttributeValueParamDetail DeleteResourceAttributeValue detail param
type DeleteResourceAttributeValueParamDetail struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
}

// DeleteResourceAttributeValueParam DeleteResourceAttributeValue request param
type DeleteResourceAttributeValueParam struct {
	BaseParam
	Params DeleteResourceAttributeValueParamDetail `json:"deleteResourceAttributeValue"`
}
// CreateResourceAttributeValueParamDetail CreateResourceAttributeValue detail param
type CreateResourceAttributeValueParamDetail struct {
	Value string `json:"value" validate:"required"`
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
}

// CreateResourceAttributeValueParam CreateResourceAttributeValue request param
type CreateResourceAttributeValueParam struct {
	BaseParam
	Params CreateResourceAttributeValueParamDetail `json:"params"`
}
