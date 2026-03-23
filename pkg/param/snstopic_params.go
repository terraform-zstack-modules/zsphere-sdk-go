// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteSNSTopicParamDetail DeleteSNSTopic detail param
type DeleteSNSTopicParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteSNSTopicParam DeleteSNSTopic request param
type DeleteSNSTopicParam struct {
	BaseParam
	Params DeleteSNSTopicParamDetail `json:"deleteSNSTopic"`
}
// CreateSNSTopicParamDetail CreateSNSTopic detail param
type CreateSNSTopicParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Locale *string `json:"locale,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSTopicParam CreateSNSTopic request param
type CreateSNSTopicParam struct {
	BaseParam
	Params CreateSNSTopicParamDetail `json:"params"`
}
// UpdateSNSTopicParamDetail UpdateSNSTopic detail param
type UpdateSNSTopicParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Locale *string `json:"locale,omitempty"`
}

// UpdateSNSTopicParam UpdateSNSTopic request param
type UpdateSNSTopicParam struct {
	BaseParam
	Params UpdateSNSTopicParamDetail `json:"updateSNSTopic"`
}
