// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreatePriceTableParamDetail CreatePriceTable detail param
type CreatePriceTableParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Prices []CreatePriceTable_PriceParam `json:"prices" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePriceTableParam CreatePriceTable request param
type CreatePriceTableParam struct {
	BaseParam
	Params CreatePriceTableParamDetail `json:"params"`
}
// UpdatePriceTableParamDetail UpdatePriceTable detail param
type UpdatePriceTableParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdatePriceTableParam UpdatePriceTable request param
type UpdatePriceTableParam struct {
	BaseParam
	Params UpdatePriceTableParamDetail `json:"updatePriceTable"`
}
// DeletePriceTableParamDetail DeletePriceTable detail param
type DeletePriceTableParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePriceTableParam DeletePriceTable request param
type DeletePriceTableParam struct {
	BaseParam
	Params DeletePriceTableParamDetail `json:"deletePriceTable"`
}
