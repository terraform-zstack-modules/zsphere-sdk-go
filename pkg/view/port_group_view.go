package view

type PortGroupInventoryView struct {
	BaseInfoView
	BaseTimeView

	VSwitchUuid     string                 `json:"vSwitchUuid"`
	VlanId          string                 `json:"vlanId"`
	VlanRanges      string                 `json:"vlanRanges"`
	VlanMode        string                 `json:"vlanMode"` //ACCESS, TRUNK, PVLAN
	Type            string                 `json:"type"`     // VirtualRouter  vrouter  SecurityGroup  Flat
	DnsDomain       string                 `json:"dnsDomain"`
	ZoneUuid        string                 `json:"zoneUuid"`
	L2NetworkUuid   string                 `json:"l2NetworkUuid"`
	State           string                 `json:"state"`
	IpVersion       int                    `json:"ipVersion"`
	System          bool                   `json:"system"`
	Category        string                 `json:"category"`
	EnableIPAM      string                 `json:"enableIPAM"`
	DNS             []string               `json:"dns"`
	IpRanges        []IpRangeInventoryView `json:"ipRanges"`
	NetworkServices []NetworkServices      `json:"networkServices"`
	HostRoute       []HostRoute            `json:"hostRoute"`
}
