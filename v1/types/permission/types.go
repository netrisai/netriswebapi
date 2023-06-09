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

package permission

const (
	SectionSERVICES                    string = "SERVICES"
	SectionINSTANCES                   string = "INSTANCES"
	SectionCIRCUITS                    string = "CIRCUITS"
	SectionACL                         string = "ACL"
	SectionACL_2                       string = "ACL_2.0"
	SectionACL_PORT_GROUPS             string = "ACL_PORT_GROUPS"
	SectionLOAD_BALANCER               string = "LOAD_BALANCER"
	SectionL4_LOAD_BALANCER            string = "L4_LOAD_BALANCER"
	SectionNET                         string = "NET"
	SectionTOPOLOGY                    string = "TOPOLOGY"
	SectionINVENTORY                   string = "INVENTORY"
	SectionSWITCH_PORTS                string = "SWITCH_PORTS"
	SectionSITES                       string = "SITES"
	SectionE_BGP                       string = "E_BGP"
	SectionE_BGP_OBJECTS               string = "E_BGP_OBJECTS"
	SectionE_BGP_ROUTE_MAPS            string = "E_BGP_ROUTE_MAPS"
	SectionSUBNETS                     string = "SUBNETS"
	SectionNAT                         string = "NAT"
	SectionVPN                         string = "VPN"
	SectionROUTES                      string = "ROUTES"
	SectionLOOKING_GLASS               string = "LOOKING_GLASS"
	SectionACCOUNTS                    string = "ACCOUNTS"
	SectionUSERS                       string = "USERS"
	SectionTENANTS                     string = "TENANTS"
	SectionUSER_ROLES                  string = "USER_ROLES"
	SectionPERMISSION_GROUPS           string = "PERMISSION_GROUPS"
	SectionGLOBAL_SETTINGS             string = "GLOBAL_SETTINGS"
	SectionGS_GENERAL                  string = "GS_GENERAL"
	SectionLOGIN_WHITELIST             string = "LOGIN_WHITELIST"
	SectionAUTHENTICATION              string = "AUTHENTICATION"
	SectionGENERAL                     string = "GENERAL"
	SectionAPI_DOCS                    string = "API_DOCS"
	SectionMONITORING_CHECK_THRESHOLDS string = "MONITORING_CHECK_THRESHOLDS"
)

type PermissionGroup struct {
	CreateDate   int        `json:"create_date"`
	Description  string     `json:"description"`
	ExternalAcl  string     `json:"external_acl"`
	Hidden       HiddenList `json:"hidden"`
	ID           int        `json:"id"`
	ModifiedDate int        `json:"modified_date"`
	Name         string     `json:"name"`
	Readonly     HiddenList `json:"readonly"`
}

type HiddenList string

type ReadonlyList string

type PermissionGroupAdd struct {
	Description string   `json:"description"`
	ExternalACL bool     `json:"externalACL"`
	Hidden      []string `json:"hidden"`
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name"`
	ReadOnly    []string `json:"readOnly"`
}
