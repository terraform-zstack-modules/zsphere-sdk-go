// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateAccessKeyParamDetail CreateAccessKey detail param
type CreateAccessKeyParamDetail struct {
	AccountUuid string `json:"accountUuid" validate:"required"`
	UserUuid *string `json:"userUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	AccessKeyID *string `json:"AccessKeyID,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAccessKeyParam CreateAccessKey request param
type CreateAccessKeyParam struct {
	BaseParam
	Params CreateAccessKeyParamDetail `json:"params"`
}
// DeleteAccessKeyParamDetail DeleteAccessKey detail param
type DeleteAccessKeyParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAccessKeyParam DeleteAccessKey request param
type DeleteAccessKeyParam struct {
	BaseParam
	Params DeleteAccessKeyParamDetail `json:"deleteAccessKey"`
}
