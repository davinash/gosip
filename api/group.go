package api

import (
	"encoding/json"
	"fmt"

	"github.com/koltyakov/gosip"
)

// Group represents SharePoint Site Groups API queryable object struct
// Always use NewGroup constructor instead of &Group{}
type Group struct {
	client    *gosip.SPClient
	config    *RequestConfig
	endpoint  string
	modifiers *ODataMods
}

// GroupInfo - site group API response payload structure
type GroupInfo struct {
	AllowMembersEditMembership     bool   `json:"AllowMembersEditMembership"`
	AllowRequestToJoinLeave        bool   `json:"AllowRequestToJoinLeave"`
	AutoAcceptRequestToJoinLeave   bool   `json:"AutoAcceptRequestToJoinLeave"`
	Description                    string `json:"Description"`
	ID                             int    `json:"Id"`
	IsHiddenInUI                   bool   `json:"IsHiddenInUI"`
	LoginName                      string `json:"LoginName"`
	OnlyAllowMembersViewMembership bool   `json:"OnlyAllowMembersViewMembership"`
	OwnerTitle                     string `json:"OwnerTitle"`
	PrincipalType                  int    `json:"PrincipalType"`
	RequestToJoinLeaveEmailSetting bool   `json:"RequestToJoinLeaveEmailSetting"`
	Title                          string `json:"Title"`
}

// GroupResp - group response type with helper processor methods
type GroupResp []byte

// NewGroup - Group struct constructor function
func NewGroup(client *gosip.SPClient, endpoint string, config *RequestConfig) *Group {
	return &Group{
		client:    client,
		endpoint:  endpoint,
		config:    config,
		modifiers: NewODataMods(),
	}
}

// ToURL gets endpoint with modificators raw URL
func (group *Group) ToURL() string {
	return toURL(group.endpoint, group.modifiers)
}

// Conf receives custom request config definition, e.g. custom headers, custom OData mod
func (group *Group) Conf(config *RequestConfig) *Group {
	group.config = config
	return group
}

// Select adds $select OData modifier
func (group *Group) Select(oDataSelect string) *Group {
	group.modifiers.AddSelect(oDataSelect)
	return group
}

// Expand adds $expand OData modifier
func (group *Group) Expand(oDataExpand string) *Group {
	group.modifiers.AddExpand(oDataExpand)
	return group
}

// Get ...
func (group *Group) Get() (GroupResp, error) {
	sp := NewHTTPClient(group.client)
	return sp.Get(group.ToURL(), getConfHeaders(group.config))
}

// Update ...
func (group *Group) Update(body []byte) ([]byte, error) {
	body = patchMetadataType(body, "SP.Group")
	sp := NewHTTPClient(group.client)
	return sp.Update(group.endpoint, body, getConfHeaders(group.config))
}

// Users ...
func (group *Group) Users() *Users {
	return NewUsers(
		group.client,
		fmt.Sprintf("%s/Users", group.endpoint),
		group.config,
	)
}

// AddUser ...
func (group *Group) AddUser(loginName string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/Users", group.ToURL())
	sp := NewHTTPClient(group.client)

	metadata := make(map[string]interface{})
	metadata["__metadata"] = map[string]string{
		"type": "SP.User",
	}
	metadata["LoginName"] = loginName
	body, _ := json.Marshal(metadata)
	return sp.Post(endpoint, body, getConfHeaders(group.config))
}

/* Response helpers */

// Data : to get typed data
func (groupResp *GroupResp) Data() *GroupInfo {
	data := parseODataItem(*groupResp)
	res := &GroupInfo{}
	json.Unmarshal(data, &res)
	return res
}

// Unmarshal : to unmarshal to custom object
func (groupResp *GroupResp) Unmarshal(obj interface{}) error {
	data := parseODataItem(*groupResp)
	return json.Unmarshal(data, obj)
}
