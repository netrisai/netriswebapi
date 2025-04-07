/*
Copyright 2023. Netris, Inc.

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

package pkeyledger

import (
	"encoding/json"
	"fmt"

	"github.com/netrisai/netriswebapi/http"
	v2address "github.com/netrisai/netriswebapi/http/addresses/v2"
)

type Client struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *Client {
	return &Client{c}
}

func parse(APIResult *http.APIResponse) ([]*Pkeyledger, error) {
	var items []*Pkeyledger
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

func parseSingle(APIResult *http.APIResponse) (*Pkeyledger, error) {
	var items *Pkeyledger
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

func (c *Client) Get() ([]*Pkeyledger, error) {
	address := c.client.URL.String() + v2address.Pkeyledger
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetPkeyledger} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetPkeyledger} %s", err)
	}
	return items, nil
}

func (c *Client) GetByUfmID(ufmID string) ([]*Pkeyledger, error) {
	address := c.client.URL.String() + v2address.Pkeyledger + fmt.Sprintf("/?ufmID=%s", ufmID)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetPkeyledgerByUfmID} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetPkeyledgerByUfmID} %s", err)
	}
	return items, nil
}

func (c *Client) GetByUfmAndPkeyID(ufmID string, pkeyID string) ([]*Pkeyledger, error) {
	address := c.client.URL.String() + v2address.Pkeyledger + fmt.Sprintf("?ufmID=%s&pkeyID=%s", ufmID, pkeyID)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetPkeyledgerByUfmAndPkeyID} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetPkeyledgerByUfmAndPkeyID} %s", err)
	}
	return items, nil
}

func (c *Client) Add(pkeyledger *PkeyledgerW) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(pkeyledger)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.Pkeyledger
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *Client) Update(ufmID string, pkeyID string, pkeyledger *PkeyledgerU) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(pkeyledger)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{UpdatePkeyledger} %s", err)
	}
	address := c.client.URL.String() + v2address.Pkeyledger + fmt.Sprintf("/%s/%s", ufmID, pkeyID)
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{UpdatePkeyledger} %s", err)
	}

	return reply, nil
}

func (c *Client) Delete(ufmID string, pkeyID string) (reply http.HTTPReply, err error) {
	address := c.client.URL.String() + v2address.Pkeyledger + fmt.Sprintf("/%s/%s", ufmID, pkeyID)
	reply, err = c.client.Delete(address, nil)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
