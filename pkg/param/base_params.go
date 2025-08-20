// Copyright (c) ZStack.io, Inc.

package param

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
