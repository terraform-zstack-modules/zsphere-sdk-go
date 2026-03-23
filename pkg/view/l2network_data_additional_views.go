// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// L2NetworkDataView L2NetworkData
type L2NetworkDataView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Description string `json:"description,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VirtualNetworkId int `json:"virtualNetworkId,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

