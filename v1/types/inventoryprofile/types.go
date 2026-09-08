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

package inventoryprofile

type Profile struct {
	CreatedDate        int                `json:"created_date"`
	CustomRules        []CustomRule       `json:"customRules"`
	Description        string             `json:"description"`
	DNSServers         string             `json:"dns_servers"`
	ID                 int                `json:"id"`
	Ipv4SSH            string             `json:"ipv4_ssh"`
	Ipv6SSH            string             `json:"ipv6_ssh"`
	ModifiedDate       int                `json:"modified_date"`
	Name               string             `json:"name"`
	NTPServers         string             `json:"ntp_servers"`
	Timezone           string             `json:"timezone"`
	FabricProps        FabricProps        `json:"fabricProps"`
	GpuClusterProps    GpuClusterProps    `json:"gpuClusterProps"`
	SNMPv2Props        SNMPv2Props        `json:"snmpv2Props"`
	ZTPProps           ZTPProps           `json:"ztpProps"`
	NetQProps          NetQProps          `json:"netqProps"`
	SyslogDestinations SyslogDestinations `json:"syslogDestinations"`
	AAAProps           AAAProps           `json:"aaa"`
}

type CustomRule struct {
	Deleted     bool   `json:"deleted,omitempty"`
	DstPort     string `json:"dstPort"`
	ID          int    `json:"id"`
	Protocol    string `json:"protocol"`
	SrcPort     string `json:"srcPort"`
	SrcSubnet   string `json:"srcSubnet"`
	Description string `json:"description"`
}

type ProfileW struct {
	CustomRules        []CustomRule       `json:"customRules"`
	Description        string             `json:"description"`
	DNSServers         string             `json:"dns_servers"`
	ID                 int                `json:"id"`
	Ipv4List           string             `json:"ipv4_list"`
	Ipv6List           string             `json:"ipv6_list"`
	Name               string             `json:"name"`
	NTPServers         string             `json:"ntp_servers"`
	Timezone           Timezone           `json:"timezone"`
	FabricProps        FabricProps        `json:"fabricProps"`
	GpuClusterProps    GpuClusterProps    `json:"gpuClusterProps"`
	SNMPv2Props        SNMPv2Props        `json:"snmpv2Props"`
	ZTPProps           ZTPProps           `json:"ztpProps"`
	NetQProps          NetQProps          `json:"netqProps"`
	SyslogDestinations SyslogDestinations `json:"syslogDestinations"`
	AAAProps           AAAProps           `json:"aaa"`
}

type NetQProps struct {
	Enabled     bool     `json:"enabled"`
	ServerAddrs []string `json:"serverAddrs"`
	ServerPort  int32    `json:"serverPort"`
}

type Timezone struct {
	Label  string `json:"label"`
	Offset string `json:"offset"`
	TzCode string `json:"tzCode"`
}

type FabricProps struct {
	OptimiseBgpOverlay           bool   `json:"optimiseBgpOverlay"`
	OptimiseBgpOverlayHypervisor bool   `json:"optimiseBgpOverlayHypervisor"`
	UnnumberedBgpUnderlay        bool   `json:"unnumberedBgpUnderlay"`
	AutomaticLinkAggregation     bool   `json:"automaticLinkAggregation"`
	MCLag                        bool   `json:"mclag"`
	ServerBasedESI               bool   `json:"serverBasedESI"`
	FabricType                   string `json:"fabricType"`
}

type GpuClusterProps struct {
	Roce                 bool   `json:"roce"`
	RoceAdaptiveRouting  bool   `json:"roceAdaptiveRouting"`
	CongestionControl    bool   `json:"congestionControl"`
	AsicMonitoring       bool   `json:"asicMonitoring"`
	Hwmp                 bool   `json:"hwmp"`
	RefArch              string `json:"refArch"`
	AggregateL3VpnPrefix bool   `json:"aggregateL3VpnPrefix"`
}

type SNMPv2Props struct {
	Enabled   bool     `json:"enabled"`
	Community string   `json:"community"`
	Ipv4List  []string `json:"ipv4_list"`
	Contact   string   `json:"contact"`
	Location  string   `json:"location"`
}

type ZTPProps struct {
	NOSImage string `json:"nosImage"`
	Password string `json:"password"`
}

type SyslogDestinations struct {
	Enabled    bool           `json:"enabled"`
	UseRfc5424 bool           `json:"useRfc5424"`
	Servers    []SyslogServer `json:"servers"`
}

type SyslogServer struct {
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	Protocol string `json:"protocol"`
	Severity string `json:"severity"`
}

type AAAProps struct {
	AuthOrder []string       `json:"authOrder"`
	Radius    RadiusProps    `json:"radius"`
	Local     LocalAuthProps `json:"local"`
}

type RadiusProps struct {
	Enabled bool `json:"enabled"`
	PriorityServers []RadiusServer `json:"priorityServers"`
}

type RadiusServer struct {
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	Priority int32  `json:"priority"`
	AuthType string `json:"authType"`
	Secret   string `json:"secret,omitempty"`
}

type LocalAuthProps struct {
	Enabled bool `json:"enabled"`
}
