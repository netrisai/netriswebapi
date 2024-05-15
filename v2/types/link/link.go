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

package link

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/netrisai/netriswebapi/http"
	v2address "github.com/netrisai/netriswebapi/http/addresses/v2"
)

type Client struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *Client {
	return &Client{c}
}

func parse(APIResult *http.APIResponse) ([]*Link, error) {
	var item []*Link
	err := http.Decode(APIResult.Data, &item)
	if err != nil {
		return item, fmt.Errorf("{parse Links} %s", err)
	}
	return item, nil
}

func parseSingle(APIResult *http.APIResponse) (*Link, error) {
	var item *Link
	err := http.Decode(APIResult.Data, &item)
	if err != nil {
		return item, fmt.Errorf("{parse Link} %s", err)
	}
	return item, nil
}

func (c *Client) Get() ([]*Link, error) {
	address := c.client.URL.String() + v2address.Links
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetLinks} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetLinks} %s", err)
	}
	return items, nil
}

func (c *Client) GetByID(id int) (*Link, error) {
	address := c.client.URL.String() + v2address.Links + "/" + strconv.Itoa(id)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetLink} %s", err)
	}

	item, err := parseSingle(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetLink} %s", err)
	}
	return item, nil
}

func (c *Client) Add(link *Linkw) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(link)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.Links
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, fmt.Errorf("{Add} %s", err)
	}

	return reply, nil
}

func (c *Client) Update(id int, link *Linkw) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(link)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.Links + "/" + strconv.Itoa(id)
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{Update} %s", err)
	}

	return reply, nil
}

func (c *Client) Delete(link *Link) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(link)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.Links
	reply, err = c.client.Delete(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *Client) DeletByID(id int) (reply http.HTTPReply, err error) {
	address := c.client.URL.String() + v2address.Links + "/" + strconv.Itoa(id)
	reply, err = c.client.Delete(address, nil)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
