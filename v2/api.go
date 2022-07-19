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

package v2

import (
	"github.com/netrisai/netriswebapi/http"
	"github.com/netrisai/netriswebapi/v1/types/acl"
	"github.com/netrisai/netriswebapi/v1/types/acl2"
	"github.com/netrisai/netriswebapi/v1/types/bgpobject"
	"github.com/netrisai/netriswebapi/v1/types/gsetting"
	"github.com/netrisai/netriswebapi/v1/types/inventoryprofile"
	"github.com/netrisai/netriswebapi/v1/types/permission"
	"github.com/netrisai/netriswebapi/v1/types/portgroup"
	"github.com/netrisai/netriswebapi/v1/types/route"
	"github.com/netrisai/netriswebapi/v1/types/routemap"
	"github.com/netrisai/netriswebapi/v1/types/site"
	"github.com/netrisai/netriswebapi/v1/types/subnet"
	"github.com/netrisai/netriswebapi/v1/types/tenant"
	"github.com/netrisai/netriswebapi/v1/types/user"
	"github.com/netrisai/netriswebapi/v1/types/userrole"
	"github.com/netrisai/netriswebapi/v2/types/bgp"
	"github.com/netrisai/netriswebapi/v2/types/inventory"
	"github.com/netrisai/netriswebapi/v2/types/ipam"
	"github.com/netrisai/netriswebapi/v2/types/l4lb"
	"github.com/netrisai/netriswebapi/v2/types/link"
	"github.com/netrisai/netriswebapi/v2/types/nat"
	"github.com/netrisai/netriswebapi/v2/types/port"
	"github.com/netrisai/netriswebapi/v2/types/roh"
	"github.com/netrisai/netriswebapi/v2/types/vnet"
)

type Clientset struct {
	Client           *http.HTTPCred
	vnet             *vnet.VNetClient
	inventory        *inventory.InventoryClient
	site             *site.SiteClient
	gsetting         *gsetting.GSettingClient
	l4lb             *l4lb.LBClient
	subnet           *subnet.SubnetClient
	port             *port.PortClient
	tenant           *tenant.TenantClient
	bgp              *bgp.BGPClient
	route            *route.RouteClient
	ipam             *ipam.IPAMClient
	user             *user.Client
	permissiongroup  *permission.Client
	userrole         *userrole.Client
	link             *link.Client
	acl              *acl.Client
	portgroup        *portgroup.Client
	roh              *roh.Client
	inventoryprofile *inventoryprofile.Client
	bgpobject        *bgpobject.Client
	routemap         *routemap.Client
	nat              *nat.Client
	acl2             *acl2.Client
}

func (c *Clientset) VNet() *vnet.VNetClient {
	if c.vnet == nil {
		c.vnet = vnet.New(c.Client)
	}
	return c.vnet
}

func (c *Clientset) Inventory() *inventory.InventoryClient {
	if c.inventory == nil {
		c.inventory = inventory.New(c.Client)
	}
	return c.inventory
}

func (c *Clientset) Site() *site.SiteClient {
	if c.site == nil {
		c.site = site.New(c.Client)
	}
	return c.site
}

func (c *Clientset) GlobalSettings() *gsetting.GSettingClient {
	if c.gsetting == nil {
		c.gsetting = gsetting.New(c.Client)
	}
	return c.gsetting
}

func (c *Clientset) L4LB() *l4lb.LBClient {
	if c.l4lb == nil {
		c.l4lb = l4lb.New(c.Client)
	}
	return c.l4lb
}

func (c *Clientset) Subnet() *subnet.SubnetClient {
	if c.subnet == nil {
		c.subnet = subnet.New(c.Client)
	}
	return c.subnet
}

func (c *Clientset) Port() *port.PortClient {
	if c.port == nil {
		c.port = port.New(c.Client)
	}
	return c.port
}

func (c *Clientset) Tenant() *tenant.TenantClient {
	if c.tenant == nil {
		c.tenant = tenant.New(c.Client)
	}
	return c.tenant
}

func (c *Clientset) BGP() *bgp.BGPClient {
	if c.bgp == nil {
		c.bgp = bgp.New(c.Client)
	}
	return c.bgp
}

func (c *Clientset) Route() *route.RouteClient {
	if c.route == nil {
		c.route = route.New(c.Client)
	}
	return c.route
}

func (c *Clientset) IPAM() *ipam.IPAMClient {
	if c.ipam == nil {
		c.ipam = ipam.New(c.Client)
	}
	return c.ipam
}

func (c *Clientset) User() *user.Client {
	if c.user == nil {
		c.user = user.New(c.Client)
	}
	return c.user
}

func (c *Clientset) Permission() *permission.Client {
	if c.permissiongroup == nil {
		c.permissiongroup = permission.New(c.Client)
	}
	return c.permissiongroup
}

func (c *Clientset) UserRole() *userrole.Client {
	if c.userrole == nil {
		c.userrole = userrole.New(c.Client)
	}
	return c.userrole
}

func (c *Clientset) Link() *link.Client {
	if c.link == nil {
		c.link = link.New(c.Client)
	}
	return c.link
}

func (c *Clientset) ACL() *acl.Client {
	if c.acl == nil {
		c.acl = acl.New(c.Client)
	}
	return c.acl
}

func (c *Clientset) PortGroup() *portgroup.Client {
	if c.portgroup == nil {
		c.portgroup = portgroup.New(c.Client)
	}
	return c.portgroup
}

func (c *Clientset) ROH() *roh.Client {
	if c.roh == nil {
		c.roh = roh.New(c.Client)
	}
	return c.roh
}

func (c *Clientset) InventoryProfile() *inventoryprofile.Client {
	if c.inventoryprofile == nil {
		c.inventoryprofile = inventoryprofile.New(c.Client)
	}
	return c.inventoryprofile
}

func (c *Clientset) BGPObject() *bgpobject.Client {
	if c.bgpobject == nil {
		c.bgpobject = bgpobject.New(c.Client)
	}
	return c.bgpobject
}

func (c *Clientset) RouteMap() *routemap.Client {
	if c.routemap == nil {
		c.routemap = routemap.New(c.Client)
	}
	return c.routemap
}

func (c *Clientset) NAT() *nat.Client {
	if c.nat == nil {
		c.nat = nat.New(c.Client)
	}
	return c.nat
}

func (c *Clientset) ACL2() *acl2.Client {
	if c.acl2 == nil {
		c.acl2 = acl2.New(c.Client)
	}
	return c.acl2
}

func Client(address, login, password string, timeout int) (*Clientset, error) {
	client, err := http.NewHTTPCredentials(address, login, password, timeout)
	if err != nil {
		return nil, err
	}
	return &Clientset{
		Client: client,
	}, nil
}

func ClientWithCookie(address, sessionID string, timeout int) (*Clientset, error) {
	client, err := http.NewHTTPCredentialsWithCooke(address, sessionID, timeout)
	if err != nil {
		return nil, err
	}
	return &Clientset{
		Client: client,
	}, nil
}
