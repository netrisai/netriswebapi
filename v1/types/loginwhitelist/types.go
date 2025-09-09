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

// LoginWhitelistEntry represents a single entry in the login whitelist (GET response format)
type LoginWhitelistEntry struct {
	ID          int    `json:"id"`
	Prefix      string `json:"prefix"`
	Length      int    `json:"length"`
	Description string `json:"description"`
}

// LoginWhitelistAdd represents the structure for adding a new entry to the login whitelist
type LoginWhitelistAdd struct {
	ID          *int   `json:"id"` // null for add operations
	IP          string `json:"ip"`
	Description string `json:"description"`
}

// LoginWhitelistUpdate represents the structure for updating an existing login whitelist entry
type LoginWhitelistUpdate struct {
	ID          int    `json:"id"`
	IP          string `json:"ip"`
	Description string `json:"description"`
}

// LoginWhitelistDelete represents the structure for deleting a login whitelist entry
type LoginWhitelistDelete struct {
	ID int `json:"id"`
}
