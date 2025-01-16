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

package l4lb

// LoadBalancer .
type LoadBalancer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Tenant    IDName `json:"tenant"`
	SiteID    int    `json:"siteId"`
	SiteName  string `json:"siteName"`
	Automatic bool   `json:"automatic"`

	KubenetInfo LoadBalancerKubenetInfo `json:"kubenet_info"`

	Protocol string `json:"protocol"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`

	Status      string            `json:"status"`
	HealthCheck LBHealthCheck     `json:"healthCheck"`
	Label       LoadBalancerLabel `json:"label"`

	BackendIPs []LBBackend `json:"backendIps"`

	CreatedDate  int    `json:"createdDate"`
	ModifiedDate int    `json:"modifiedDate"`
	Vpc          IDName `json:"vpc"`
}

// LoadBalancerLabel .
type LoadBalancerLabel struct {
	Status string `json:"status"`
	Text   string `json:"text"`
}

// LoadBalancerAddResponse .
type LoadBalancerAddResponse struct {
	ID int    `json:"id"`
	IP string `json:"ip"`
}

// LoadBalancerKubenetInfo .
type LoadBalancerKubenetInfo struct {
	KubenetID  int    `json:"kubenet_id"`
	ServiceUID string `json:"service_uid"`
	Namespace  string `json:"namespace"`
}

// LoadBalancerAdd .
type LoadBalancerAdd struct {
	Name string `json:"name"`

	Tenant    IDName `json:"tenant"`
	SiteID    int    `json:"siteId"`
	Automatic bool   `json:"automatic"`
	Internal  int    `json:"internal"`

	KubenetInfoString string `json:"kubenet_info"`

	Protocol string `json:"protocol"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`

	Status string `json:"status"`

	HealthCheck string `json:"healthCheck"`
	Timeout     string `json:"timeOut"`
	RequestPath string `json:"requestPath"`

	Backend []LBAddBackend `json:"backendIps"`
	Vpc          IDName `json:"vpc"`
}

// LoadBalancerUpdate .
type LoadBalancerUpdate struct {
	Name      string `json:"name"`
	Tenant    IDName `json:"tenant"`
	SiteID    int    `json:"siteId"`
	SiteName  string `json:"siteName"`
	Automatic bool   `json:"automatic"`

	Protocol string `json:"protocol"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	Status   string `json:"status"`

	KubenetInfoString string `json:"kubenet_info"`

	HealthCheck string `json:"healthCheck"`
	Timeout     string `json:"timeOut"`
	RequestPath string `json:"requestPath"`

	BackendIPs []LBBackend `json:"backendIps"`
	Vpc        IDName      `json:"vpc"`
}

// LBBackend .
type LBAddBackend struct {
	IP          string `json:"ip"`
	Port        int    `json:"port"`
	Maintenance bool   `json:"maintenance"`
}

// LBBackend .
type LBBackend struct {
	ID          string `json:"id"`
	IP          string `json:"ip"`
	Port        string `json:"port"`
	Status      string `json:"status"`
	Response    string `json:"response"`
	Maintenance bool   `json:"maintenance"`
}

// LBHealthCheckTCP .
type LBHealthCheckTCP struct {
	Timeout     string `json:"timeOut"`
	RequestPath string `json:"requestPath"`
}

// LBHealthCheckHTTP .
type LBHealthCheckHTTP struct {
	Timeout     string `json:"timeOut"`
	RequestPath string `json:"requestPath"`
}

// LBHealthCheck .
type LBHealthCheck struct {
	TCP  LBHealthCheckTCP
	HTTP LBHealthCheckHTTP
}

type IDName struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
