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

package loginwhitelist

import (
	"encoding/json"
	"fmt"

	"github.com/netrisai/netriswebapi/http"
	v1address "github.com/netrisai/netriswebapi/http/addresses/v1"
)

type LoginWhitelistClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *LoginWhitelistClient {
	return &LoginWhitelistClient{c}
}

// Get retrieves all login whitelist entries
func (c *LoginWhitelistClient) Get() ([]*LoginWhitelistEntry, error) {
	address := c.client.URL.String() + v1address.LoginWhitelist
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetLoginWhitelist} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetLoginWhitelist} %s", err)
	}
	return items, nil
}

// Add creates a new login whitelist entry
func (c *LoginWhitelistClient) Add(entry *LoginWhitelistAdd) (reply http.HTTPReply, err error) {
	// Set ID to null for add operations
	entry.ID = nil
	js, err := json.Marshal(entry)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.LoginWhitelist
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

// Update modifies an existing login whitelist entry
func (c *LoginWhitelistClient) Update(entry *LoginWhitelistUpdate) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(entry)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{UpdateLoginWhitelist} %s", err)
	}
	address := c.client.URL.String() + v1address.LoginWhitelist
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{UpdateLoginWhitelist} %s", err)
	}

	return reply, nil
}

// Delete removes a login whitelist entry
func (c *LoginWhitelistClient) Delete(id int) (reply http.HTTPReply, err error) {
	deleteEntry := &LoginWhitelistDelete{ID: id}
	js, err := json.Marshal(deleteEntry)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.LoginWhitelist
	reply, err = c.client.Delete(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

// parse parses the API response into LoginWhitelistEntry structs
func parse(APIResult *http.APIResponse) ([]*LoginWhitelistEntry, error) {
	var items []*LoginWhitelistEntry
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}
