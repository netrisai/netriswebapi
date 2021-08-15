package v2

import (
	"github.com/netrisai/netriswebapi/http"
	"github.com/netrisai/netriswebapi/v2/types/vnet"
)

type Clientset struct {
	Client *http.HTTPCred
	vnet   *vnet.VNetClient
}

func (c *Clientset) VNet() *vnet.VNetClient {
	if c.vnet == nil {
		c.vnet = vnet.New(c.Client)
	}
	return c.vnet
}

func Client(address, login, password string, timeout int) (*Clientset, error) {
	client, err := http.NewHTTPCredentials(address, login, password, timeout)
	if err != nil {
		return nil, err
	}
	if err := client.LoginUser(); err != nil {
		return nil, err
	}
	return &Clientset{
		Client: client,
	}, nil
}
