// Copyright (c) ZStack.io, Inc.

package param

type CreateL2VirtualSwitchParam struct {
	BaseParam
	Params CreateCreateL2VirtualSwitchDetailParam `json:"params"`
}

type CreateCreateL2VirtualSwitchDetailParam struct {
	IsDistributed bool `json:"isDistributed"`

	Name              string `json:"name"`
	Description       string `json:"description"`
	ZoneUuid          string `json:"zoneUuid"`          // Zone UUID
	PhysicalInterface string `json:"physicalInterface"` // Physical network interface
	Type              string `json:"type"`              // Layer 2 network type
	VSwitchType       string `json:"vSwitchType"`       // VIRTUAL_MACHINE, VIRTUAL_SWITCH
	ResourceUuid      string `json:"resourceUuid"`      // Resource UUID
}
