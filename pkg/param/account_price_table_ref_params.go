// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// GetAccountPriceTableRefParamDetail GetAccountPriceTableRef detail param
type GetAccountPriceTableRefParamDetail struct {
	TableUuid *string `json:"tableUuid,omitempty"`
	AccountUuid *string `json:"accountUuid,omitempty"`
}

// GetAccountPriceTableRefParam GetAccountPriceTableRef request param
type GetAccountPriceTableRefParam struct {
	BaseParam
	Params GetAccountPriceTableRefParamDetail `json:"getAccountPriceTableRef"`
}
