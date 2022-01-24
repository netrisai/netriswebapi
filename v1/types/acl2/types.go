/*
Copyright 2021. Netris, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package acl2

type ACL2 struct {
	ExternalAcl       bool              `json:"externalAcl"`
	ID                int               `json:"id"`
	Lbs               []interface{}     `json:"lbs"`
	Name              string            `json:"name"`
	Prefixes          []Prefix          `json:"prefixes"`
	Privacy           string            `json:"privacy"`
	Protoports        []Protoport       `json:"protoports"`
	PubInstances      []PubInstance     `json:"pubInstances"`
	PublisherPrefixes []PublisherPrefix `json:"publisherPrefixes"`
	ServiceName       string            `json:"serviceName"`
	Status            string            `json:"status"`
	SubInstances      []SubInstance     `json:"subInstances"`
	TenantID          int               `json:"tenantID"`
	TenantName        string            `json:"tenantName"`
}

type Prefix struct {
	Approve  []interface{} `json:"approve"`
	Comment  string        `json:"comment"`
	ID       int           `json:"id"`
	Length   int           `json:"length"`
	Prefix   string        `json:"prefix"`
	TenantID int           `json:"tenantID"`
	Value    string        `json:"value"`
}

type Protoport struct {
	Description   string `json:"description"`
	ID            int    `json:"id"`
	Port          string `json:"port"`
	PortGroupID   string `json:"portGroupID"`
	PortGroupName string `json:"portGroupName"`
	Proto         string `json:"proto"`
}

type PubInstance struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type PublisherPrefix struct {
	Approve   []interface{} `json:"approve"`
	ID        int           `json:"id"`
	Length    string        `json:"length"`
	Prefix    string        `json:"prefix"`
	ServiceID int           `json:"serviceID"`
	Value     string        `json:"value"`
}

type SubInstance struct {
	Approve    []interface{} `json:"approve"`
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	TenantID   int           `json:"tenantID"`
	TenantName string        `json:"tenantName"`
	Type       string        `json:"type"`
}

type ACLw struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name"`
	Privacy  string `json:"privacy"`
	TenantID int    `json:"tenantID"`
}

type ACLStatusW struct {
	ID       int    `json:"id"`
	Status   string `json:"status"`
	TenantID int    `json:"tenantID"`
}

type PublisherW struct {
	ID        int                  `json:"id"`
	Instances []int                `json:"instances,omitempty"`
	Lbs       []PublisherWLB       `json:"lbs"`
	Prefixes  []PublisherWPrefix   `json:"prefixes,omitempty"`
	Protocols []PublisherWProtocol `json:"protocols,omitempty"`
	TenantID  int                  `json:"tenantID"`
	Type      string               `json:"type"`
}

type PublisherWLB struct {
	ID        int    `json:"id"`
	IPAddress string `json:"ipAddress"`
}

type PublisherWPrefix struct {
	ID     int    `json:"id"`
	Length string `json:"length"`
	Prefix string `json:"prefix"`
}

type PublisherD struct {
	ID        int    `json:"id"`
	Length    string `json:"length,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	ServiceID int    `json:"serviceID"`
	Type      string `json:"type"`
}

type PublisherWProtocol struct {
	Description string `json:"description"`
	ID          int    `json:"id"`
	Port        string `json:"port"`
	PortGroupID int    `json:"portGroupID,omitempty"`
	Proto       string `json:"proto"`
}

type PublisherApprove struct {
	Action    string `json:"action"`
	ApproveID int    `json:"approveID"`
	PrefixID  int    `json:"prefixID"`
}

type PublisherReject PublisherApprove

type SubscriberW struct {
	ID        int                 `json:"id"`
	Instances []int               `json:"instances"`
	Prefixes  []SubscriberWPrefix `json:"prefixes"`
	TenantID  int                 `json:"tenantID"`
	Type      string              `json:"type"`
}

type SubscriberWPrefix struct {
	ID        int    `json:"id"`
	Comment string `json:"comment"`
	Length  string `json:"length"`
	Prefix  string `json:"prefix"`
}

type SubscriberD struct {
	ID        int    `json:"id"`
	ServiceID int    `json:"serviceID"`
	Type      string `json:"type"`
}

type SubscriberApprove struct {
	Action          string `json:"action"`
	ApproveID       int    `json:"approveID"`
	ID              int    `json:"id"`
	PrefixID        int    `json:"prefixID"`
	ServiceTenantID int    `json:"serviceTenantID"`
	SubTenantID     int    `json:"subTenantID"`
	Type            string `json:"type"`
}

type SubscriberReject SubscriberApprove
