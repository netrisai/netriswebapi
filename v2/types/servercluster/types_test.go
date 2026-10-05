package servercluster

import (
	"encoding/json"
	"testing"
)

type templateJSON struct {
	ID           int                 `json:"id"`
	IPv4Gateways []VNetGatewayAssign `json:"ipv4Gateways"`
	IPv6Gateways []VNetGatewayAssign `json:"ipv6Gateways"`
}

func marshalTemplate(t *testing.T, w ServerClusterW) (map[string]json.RawMessage, templateJSON, []byte) {
	t.Helper()
	b, err := json.Marshal(w)
	if err != nil {
		t.Fatal(err)
	}
	var top map[string]json.RawMessage
	if err := json.Unmarshal(b, &top); err != nil {
		t.Fatal(err)
	}
	var tpl templateJSON
	if err := json.Unmarshal(top["srvClusterTemplate"], &tpl); err != nil {
		t.Fatal(err)
	}
	return top, tpl, b
}

func TestServerClusterWMarshalGateways(t *testing.T) {
	v4 := VNetGatewayAssign{Postfix: "data", Address: "172.30.64.1/24"}
	v6 := VNetGatewayAssign{Postfix: "data", Address: "2001:db8::1/64"}

	top, tpl, b := marshalTemplate(t, ServerClusterW{
		Name:                           "TENANT-1",
		SrvClusterTemplate:             IDName{ID: 7},
		SrvClusterTemplateIPv4Gateways: []VNetGatewayAssign{v4},
		SrvClusterTemplateIPv6Gateways: []VNetGatewayAssign{v6},
	})

	if tpl.ID != 7 {
		t.Errorf("unexpected template id: %s", b)
	}
	if len(tpl.IPv4Gateways) != 1 || tpl.IPv4Gateways[0] != v4 {
		t.Errorf("unexpected ipv4Gateways: %s", b)
	}
	if len(tpl.IPv6Gateways) != 1 || tpl.IPv6Gateways[0] != v6 {
		t.Errorf("unexpected ipv6Gateways: %s", b)
	}
	for _, k := range []string{"SrvClusterTemplateIPv4Gateways", "SrvClusterTemplateIPv6Gateways", "ipv4Gateways", "ipv6Gateways"} {
		if _, ok := top[k]; ok {
			t.Errorf("%q must not appear at the top level: %s", k, b)
		}
	}
}

func TestServerClusterWMarshalOnlyIPv4Gateways(t *testing.T) {
	_, tpl, b := marshalTemplate(t, ServerClusterW{
		SrvClusterTemplate:             IDName{ID: 7},
		SrvClusterTemplateIPv4Gateways: []VNetGatewayAssign{{Postfix: "data", Address: "172.30.64.1/24"}},
	})

	if len(tpl.IPv4Gateways) != 1 {
		t.Errorf("unexpected ipv4Gateways: %s", b)
	}
	if len(tpl.IPv6Gateways) != 0 {
		t.Errorf("ipv6Gateways should be omitted when empty: %s", b)
	}
}

func TestServerClusterWMarshalNoGatewaysUnchanged(t *testing.T) {
	_, _, b := marshalTemplate(t, ServerClusterW{
		SrvClusterTemplate: IDName{ID: 7, Name: "tpl"},
	})

	var raw struct {
		SrvClusterTemplate map[string]json.RawMessage `json:"srvClusterTemplate"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"vlans", "ipv4Gateways", "ipv6Gateways"} {
		if _, ok := raw.SrvClusterTemplate[k]; ok {
			t.Errorf("%q should be omitted when unset: %s", k, b)
		}
	}
}
