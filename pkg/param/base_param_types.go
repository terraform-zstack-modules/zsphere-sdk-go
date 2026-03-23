// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

type DeleteMode string

const (
	DeleteModePermissive DeleteMode = "Permissive"
	DeleteModeEnforcing  DeleteMode = "Enforcing"
)

type BaseParam struct {
	SystemTags []string `json:"systemTags,omitempty"` // System tags
	UserTags   []string `json:"userTags,omitempty"`   // User tags
	RequestIp  string   `json:"requestIp,omitempty"`  // Request IP
}

type HqlParam struct {
	OperationName string    `json:"operationName"` // Request name
	Query         string    `json:"query"`         // Query statement
	Variables     Variables `json:"variables"`     // Parameters for the statement
}

type Variables struct {
	Conditions      []Condition            `json:"conditions"`      // Conditions
	ExtraConditions []Condition            `json:"extraConditions"` // Extra conditions
	Input           map[string]interface{} `json:"input"`           // Input parameters
	PageVar         `json:",inline,omitempty"`
	Type            string `json:"type"` // Type
}

type Condition struct {
	Key   string `json:"key"`   // Key
	Op    string `json:"op"`    // Operator
	Value string `json:"value"` // Value
}

type PageVar struct {
	Start int `json:"start,omitempty"` // Start page
	Limit int `json:"limit,omitempty"` // Limit per page
}

// ========== Dynamically Generated Param Types ==========

// SetVmNicSecurityGroup_VmNicSecurityGroupRefAOParam VmNicSecurityGroupRefAO param struct
type SetVmNicSecurityGroup_VmNicSecurityGroupRefAOParam struct {
	SecurityGroupUuid string `json:"securityGroupUuid,omitempty"`
	Priority int `json:"priority,omitempty"`
}

// LabelParam Label param struct
type LabelParam struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
	Op string `json:"op,omitempty"`
	Compatible bool `json:"compatible,omitempty"`
}

// UpdateZStoneHostConfig_ZStoneHostAOParam ZStoneHostAO param struct
type UpdateZStoneHostConfig_ZStoneHostAOParam struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	ManagementIp *string `json:"managementIp,omitempty"`
	PublicIp *string `json:"publicIp,omitempty"`
	ClusterIp *string `json:"clusterIp,omitempty"`
	ManagementNode bool `json:"managementNode,omitempty"`
}

// AddSecurityGroupRule_SecurityGroupRuleAOParam SecurityGroupRuleAO param struct
type AddSecurityGroupRule_SecurityGroupRuleAOParam struct {
	Type string `json:"type,omitempty"`
	State *string `json:"state,omitempty"`
	Description *string `json:"description,omitempty"`
	RemoteSecurityGroupUuid *string `json:"remoteSecurityGroupUuid,omitempty"`
	IpVersion *int `json:"ipVersion,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	SrcIpRange *string `json:"srcIpRange,omitempty"`
	DstIpRange *string `json:"dstIpRange,omitempty"`
	DstPortRange *string `json:"dstPortRange,omitempty"`
	Action *string `json:"action,omitempty"`
	StartPort *int `json:"startPort,omitempty"`
	EndPort *int `json:"endPort,omitempty"`
	AllowedCidr *string `json:"allowedCidr,omitempty"`
}

// VmCustomSpecificationStructParam VmCustomSpecificationStruct param struct
type VmCustomSpecificationStructParam struct {
	Uuid string `json:"uuid,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	RootPassword *string `json:"rootPassword,omitempty"`
	GenerateSID *bool `json:"generateSID,omitempty"`
	DomainMode string `json:"domainMode,omitempty"`
	DomainName *string `json:"domainName,omitempty"`
	DomainUsername *string `json:"domainUsername,omitempty"`
	DomainPassword *string `json:"domainPassword,omitempty"`
	Organization *string `json:"organization,omitempty"`
}

// MetricDatumParam MetricDatum param struct
type MetricDatumParam struct {
	MetricName *string `json:"metricName,omitempty"`
	Value *float64 `json:"value,omitempty"`
	Time *int64 `json:"time,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}

// HostSshParameterParam HostSshParameter param struct
type HostSshParameterParam struct {
	Ip *string `json:"ip,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Port int `json:"port,omitempty"`
}

// CreateAlarm_ActionParamParam ActionParam param struct
type CreateAlarm_ActionParamParam struct {
	ActionUuid *string `json:"actionUuid,omitempty"`
	ActionType *string `json:"actionType,omitempty"`
}

// ThresholdParam Threshold param struct
type ThresholdParam struct {
	ThresholdName *string `json:"thresholdName,omitempty"`
	ThresholdValue *string `json:"thresholdValue,omitempty"`
	Operator *string `json:"operator,omitempty"`
}

// CreatePriceTable_PriceParam Price param struct
type CreatePriceTable_PriceParam struct {
	ResourceName *string `json:"resourceName,omitempty"`
	ResourceUnit *string `json:"resourceUnit,omitempty"`
	TimeUnit *string `json:"timeUnit,omitempty"`
	Price float64 `json:"price,omitempty"`
	DateInLong *int64 `json:"dateInLong,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
}

// MiniHostInfoParam MiniHostInfo param struct
type MiniHostInfoParam struct {
	Sn *string `json:"sn,omitempty"`
	DnsAddresses []string `json:"dnsAddresses,omitempty"`
	Ipmi MiniNetworkConfigStructParam `json:"ipmi,omitempty"`
	Mgmt MiniNetworkConfigStructParam `json:"mgmt,omitempty"`
}

// UpdateResourceConfigs_ResourceConfigAOParam ResourceConfigAO param struct
type UpdateResourceConfigs_ResourceConfigAOParam struct {
	Category *string `json:"category,omitempty"`
	Name string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// DiskAOParam DiskAO param struct
type DiskAOParam struct {
	Boot bool `json:"boot,omitempty"`
	Platform *string `json:"platform,omitempty"`
	GuestOsType *string `json:"guestOsType,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	Size int64 `json:"size,omitempty"`
	TemplateUuid *string `json:"templateUuid,omitempty"`
	DiskOfferingUuid *string `json:"diskOfferingUuid,omitempty"`
	SourceType *string `json:"sourceType,omitempty"`
	SourceUuid *string `json:"sourceUuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
	Name string `json:"name,omitempty"`
}

// UpdateSecurityGroupRulePriority_SecurityGroupRulePriorityAOParam SecurityGroupRulePriorityAO param struct
type UpdateSecurityGroupRulePriority_SecurityGroupRulePriorityAOParam struct {
	RuleUuid string `json:"ruleUuid,omitempty"`
	Priority int `json:"priority,omitempty"`
}

// ResourceAttributeConstraintParamParam ResourceAttributeConstraintParam param struct
type ResourceAttributeConstraintParamParam struct {
	Id int64 `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Parameter *string `json:"parameter,omitempty"`
}

// HostKernelInterfaceStructParam HostKernelInterfaceStruct param struct
type HostKernelInterfaceStructParam struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	Ip6 *string `json:"ip6,omitempty"`
	Ipv6Prefix *int `json:"ipv6Prefix,omitempty"`
}

// MiniNetworkConfigStructParam MiniNetworkConfigStruct param struct
type MiniNetworkConfigStructParam struct {
	Gw *string `json:"gw,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Vlan *string `json:"vlan,omitempty"`
	Bond *string `json:"bond,omitempty"`
}

