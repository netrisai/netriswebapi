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

package bgpobject

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

func parse(APIResult *http.APIResponse) ([]*BGPObject, error) {
	var item []*BGPObject
	err := http.Decode(APIResult.Data, &item)
	if err != nil {
		return item, fmt.Errorf("{parse} %s", err)
	}
	return item, nil
}

func (c *Client) Get() ([]*BGPObject, error) {
	address := c.client.URL.String() + v1address.BGPObjects
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetBGPObjects} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetBGPObjects} %s", err)
	}
	return items, nil
}

func (c *Client) Add(bgpObject *BGPObjectW) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(bgpObject)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.BGPObjects
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *Client) Update(bgpObject *BGPObjectW) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(bgpObject)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{Update BGP Objects} %s", err)
	}
	address := c.client.URL.String() + v1address.BGPObjects
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{Update BGP Objects} %s", err)
	}

	return reply, nil
}

func (c *Client) Delete(id int) (reply http.HTTPReply, err error) {
	lb := struct {
		ID int `json:"id"`
	}{id}
	js, err := json.Marshal(lb)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.BGPObjects
	reply, err = c.client.Delete(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
