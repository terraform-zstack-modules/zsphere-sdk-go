// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateAccessControlListParamDetail CreateAccessControlList detail param
type CreateAccessControlListParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	IpVersion *int `json:"ipVersion,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAccessControlListParam CreateAccessControlList request param
type CreateAccessControlListParam struct {
	BaseParam
	Params CreateAccessControlListParamDetail `json:"params"`
}
// DeleteAccessControlListParamDetail DeleteAccessControlList detail param
type DeleteAccessControlListParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAccessControlListParam DeleteAccessControlList request param
type DeleteAccessControlListParam struct {
	BaseParam
	Params DeleteAccessControlListParamDetail `json:"deleteAccessControlList"`
}
