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

import (
	"encoding/json"
	"fmt"

	"github.com/netrisai/netriswebapi/http"
	v1address "github.com/netrisai/netriswebapi/http/addresses/v1"
)

type Client struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *Client {
	return &Client{c}
}

func parse(APIResult *http.APIResponse) ([]*ACL2, error) {
	var sites []*ACL2
	err := http.Decode(APIResult.Data, &sites)
	if err != nil {
		return sites, fmt.Errorf("{parse} %s", err)
	}
	return sites, nil
}

func (c *Client) Get() ([]*ACL2, error) {
	address := c.client.URL.String() + v1address.ACL2
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetACL2.0s} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetACL2.0s} %s", err)
	}
	return items, nil
}

func (c *Client) Add(acl *ACLw) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(acl)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.ACL2
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *Client) Update(acl *ACLw) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(acl)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{UpdateACL2.0} %s", err)
	}
	address := c.client.URL.String() + v1address.ACL2
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{UpdateACL2.0} %s", err)
	}

	return reply, nil
}

func (c *Client) Delete(id int) (reply http.HTTPReply, err error) {
	lb := struct {
		ID []int `json:"id"`
	}{[]int{id}}
	js, err := json.Marshal(lb)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.ACL2
	reply, err = c.client.Delete(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *Client) ChangeStatus(status *ACLStatusW) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(status)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{Change ACL2.0 Status} %s", err)
	}
	address := c.client.URL.String() + v1address.ACL2Status
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{Change ACL2.0 Status} %s", err)
	}

	return reply, nil
}

func (c *Client) AddPublisher(pub *PublisherW) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(pub)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{Add ACL2.0 Publishers} %s", err)
	}
	address := c.client.URL.String() + v1address.ACL2Publishers
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{Add ACL2.0 Publishers} %s", err)
	}

	return reply, nil
}
