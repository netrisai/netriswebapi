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
	"github.com/netrisai/netriswebapi/v1/types/gsetting"
	"github.com/netrisai/netriswebapi/v1/types/site"
	"github.com/netrisai/netriswebapi/v2/types/inventory"
	"github.com/netrisai/netriswebapi/v2/types/vnet"
)

type Clientset struct {
	Client    *http.HTTPCred
	vnet      *vnet.VNetClient
	inventory *inventory.InventoryClient
	site      *site.SiteClient
	gsetting  *gsetting.GSettingClient
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

func Client(address, login, password string, timeout int) (*Clientset, error) {
	client, err := http.NewHTTPCredentials(address, login, password, timeout)
	if err != nil {
		return nil, err
	}
	if err := client.LoginUser(); err != nil {
		return nil, err
	}
	return &Clientset{
		Client: client,
	}, nil
}
