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

package tenant

import (
	"encoding/json"
	"fmt"

	"github.com/netrisai/netriswebapi/http"
	v1address "github.com/netrisai/netriswebapi/http/addresses/v1"
)

type TenantClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *TenantClient {
	return &TenantClient{c}
}

func parse(APIResult *http.APIResponse) ([]*Tenant, error) {
	var sites []*Tenant
	err := http.Decode(APIResult.Data, &sites)
	if err != nil {
		return sites, fmt.Errorf("{parse} %s", err)
	}
	return sites, nil
}

func (c *TenantClient) Get() ([]*Tenant, error) {
	address := c.client.URL.String() + v1address.Tenants
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetTenats} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetTenats} %s", err)
	}
	return items, nil
}

func (c *TenantClient) Add(tenant *Tenant) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(tenant)
	if err != nil {
		return reply, fmt.Errorf("{Tenant Add} %s", err)
	}

	address := c.client.URL.String() + v1address.Tenants
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, fmt.Errorf("{Tenant Add} %s", err)
	}

	return reply, nil
}

func (c *TenantClient) Update(tenant *Tenant) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(tenant)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{Tenant Update} %s", err)
	}
	address := c.client.URL.String() + v1address.Tenants
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{Tenant Update} %s", err)
	}

	return reply, nil
}

func (c *TenantClient) Delete(id int) (reply http.HTTPReply, err error) {
	lb := struct {
		ID int `json:"id"`
	}{id}
	js, err := json.Marshal(lb)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.Tenants
	reply, err = c.client.Delete(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
