package vnet

/*
VNet Structure for GET requests
*/

type VNet struct {
	CreatedDate  int64         `json:"createdDate"`
	Gateways     []VNetGateway `json:"gateways"`
	ID           int64         `json:"id"`
	Internal     int64         `json:"internal"`
	MacAddress   string        `json:"macAddress"`
	ModifiedDate int64         `json:"modifiedDate"`
	Name         string        `json:"name"`
	NativeVlan   int64         `json:"nativeVlan"`
	PortsCount   int64         `json:"portsCount"`
	Provisioning bool          `json:"provisioning"`
	State        string        `json:"state"`
	Status       VNetStatus    `json:"status"`
	VlanAware    bool          `json:"vlanAware"`
	Vlans        string        `json:"vlans"`
	VxlanID      int64         `json:"vxlanID"`
}

type VNetGateway struct {
	IPFamily string `json:"ipFamily"`
	Prefix   string `json:"prefix"`
	Vlan     int64  `json:"vlan"`
}

type VNetStatus struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

/*
VNet Structure for GET by id request
*/

type VNetDetailed struct {
	CreatedDate  int64                     `json:"createdDate"`
	Gateways     []VNetDetailedGateway     `json:"gateways"`
	GuestTenants []VNetDetailedGuestTenant `json:"guestTenants"`
	ID           int64                     `json:"id"`
	Internal     int64                     `json:"internal"`
	MacAddress   string                    `json:"macAddress"`
	ModifiedDate int64                     `json:"modifiedDate"`
	Name         string                    `json:"name"`
	NativeVlan   int64                     `json:"nativeVlan"`
	Ports        []VNetDetailedPort        `json:"ports"`
	PortsCount   int64                     `json:"portsCount"`
	Provisioning bool                      `json:"provisioning"`
	Sites        []VNetDetailedSite        `json:"sites"`
	State        string                    `json:"state"`
	Status       VNetStatus                `json:"status"`
	Tenant       VNetDetailedTenant        `json:"tenant"`
	VlanAware    bool                      `json:"vlanAware"`
	Vlans        string                    `json:"vlans"`
	VxlanID      int64                     `json:"vxlanID"`
}

type VNetDetailedGateway struct {
	Prefix string `json:"prefix"`
}

type VNetDetailedGuestTenant struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type VNetDetailedSite struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type VNetDetailedTenant struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type VNetDetailedPort struct {
	Access           bool                             `json:"access"`
	AdminDown        string                           `json:"adminDown"`
	AutoNeg          string                           `json:"autoNeg"`
	Breakout         int64                            `json:"breakout"`
	CreatedDate      int64                            `json:"createdDate"`
	Description      string                           `json:"description"`
	DesiredSpeed     interface{}                      `json:"desiredSpeed"`
	Duplex           string                           `json:"duplex"`
	Extension        int64                            `json:"extension"`
	ID               int64                            `json:"id"`
	IfName           interface{}                      `json:"ifName"`
	Lacp             string                           `json:"lacp"`
	MacCount         int64                            `json:"macCount"`
	ModifiedDate     int64                            `json:"modifiedDate"`
	Mtu              int64                            `json:"mtu"`
	Name             string                           `json:"name"`
	ParentPort       int64                            `json:"parentPort"`
	Port             string                           `json:"port"`
	Site             VNetDetailedPortSite             `json:"site"`
	SlavePorts       []interface{}                    `json:"slavePorts"`
	Speed            string                           `json:"speed"`
	State            VNetDetailedPortState            `json:"state"`
	StateInHierarchy VNetDetailedPortStateInHierarchy `json:"stateInHierarchy"`
	Status           VNetStatus                       `json:"status"`
	Switch           VNetDetailedPortSwitch           `json:"switch"`
	SwitchName       string                           `json:"switchName"`
	Tenant           VNetDetailedPortTenant           `json:"tenant"`
	Transceiver      string                           `json:"transceiver"`
	Untagged         bool                             `json:"untagged"`
	Used             bool                             `json:"used"`
	Vlan             string                           `json:"vlan"`
}

type VNetDetailedPortSite struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type VNetDetailedPortState struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}

type VNetDetailedPortStateInHierarchy struct {
	Aggregated       bool  `json:"aggregated"`
	Breakout         bool  `json:"breakout"`
	BreakoutChild    int64 `json:"breakoutChild"`
	Extended         int64 `json:"extended"`
	ExtensionsParent int64 `json:"extensionsParent"`
	LagMember        bool  `json:"lagMember"`
}

type VNetDetailedPortSwitch struct {
	ID          int64  `json:"id"`
	MainAddress string `json:"mainAddress"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type VNetDetailedPortTenant struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

/*
VNet Structure for POST requests
*/

type VNetAdd struct {
	Name         string           `json:"name"`
	Sites        []VNetAddSite    `json:"sites"`
	Gateways     []VNetAddGateway `json:"gateways"`
	GuestTenants []VNetAddTenant  `json:"guestTenants"`
	NativeVlan   int64            `json:"nativeVlan"`
	Ports        []VNetAddPort    `json:"ports"`
	Provisioning bool             `json:"provisioning"`
	State        string           `json:"state"`
	Tenant       VNetAddTenant    `json:"tenant"`
	VlanAware    bool             `json:"vlanAware"`
	Vlans        string           `json:"vlans"`
}

type VNetAddGateway struct {
	Prefix string `json:"prefix"`
	Vlan   int64  `json:"vlan"`
}

type VNetAddTenant struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type VNetAddPort struct {
	Access bool   `json:"access"`
	ID     int64  `json:"id"`
	Lacp   string `json:"lacp"`
	State  string `json:"state"`
	Vlan   string `json:"vlan"`
}

type VNetAddSite struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

/*
VNet Structure for PUT requests
*/

type VNetUpdate struct {
	Gateways     []VNetUpdateGateway     `json:"gateways"`
	GuestTenants []VNetUpdateGuestTenant `json:"guestTenants"`
	Name         string                  `json:"name"`
	NativeVlan   int64                   `json:"nativeVlan"`
	Ports        []VNetUpdatePort        `json:"ports"`
	Provisioning bool                    `json:"provisioning"`
	Sites        []VNetUpdateSite        `json:"sites"`
	State        string                  `json:"state"`
	Vlans        string                  `json:"vlans"`
}

type VNetUpdateGateway struct {
	Prefix string `json:"prefix"`
	Vlan   int64  `json:"vlan"`
}

type VNetUpdateGuestTenant struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type VNetUpdatePort struct {
	Access bool   `json:"access"`
	ID     int64  `json:"id"`
	Lacp   string `json:"lacp"`
	State  string `json:"state"`
	Vlan   string `json:"vlan"`
}

type VNetUpdateSite struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
