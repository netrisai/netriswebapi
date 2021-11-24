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

package routemap

type RouteMap struct {
	CreatedDate  int        `json:"createdDate"`
	ID           int        `json:"id"`
	ModifiedDate int        `json:"modifiedDate"`
	Name         string     `json:"name"`
	Sequences    []Sequence `json:"sequences"`
}

type Sequence struct {
	Actions     []SequenceAction `json:"actions"`
	Description string           `json:"description"`
	ID          int              `json:"id"`
	Matches     []SequenceMatch  `json:"matches"`
	Number      int              `json:"number"`
	Policy      string           `json:"policy"`
}

type SequenceAction struct {
	ID         string `json:"id"`
	Parameter  string `json:"parameter"`
	SequenceID string `json:"sequenceID"`
	Type       string `json:"type"`
	Value      string `json:"value"`
}

type SequenceMatch struct {
	EbgpObject     string `json:"ebgpObject"`
	EbgpObjectType string `json:"ebgpObjectType"`
	ID             string `json:"id"`
	SequenceID     string `json:"sequenceID"`
	Type           string `json:"type"`
	Value          string `json:"value"`
}
