// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateQuotaParamDetail UpdateQuota detail param
type UpdateQuotaParamDetail struct {
	IdentityUuid string `json:"identityUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Value int64 `json:"value" validate:"required"`
}

// UpdateQuotaParam UpdateQuota request param
type UpdateQuotaParam struct {
	BaseParam
	Params UpdateQuotaParamDetail `json:"updateQuota"`
}
