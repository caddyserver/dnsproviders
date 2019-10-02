// Package dyn adapts the lego Dyn DNS provider
// for Caddy. Importing this package plugs it in.
package dyn

import (
	"errors"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/go-acme/lego/v3/providers/dns/dyn"
)

func init() {
	caddytls.RegisterDNSProvider("dyn", NewDNSProvider)
}

// NewDNSProvider returns a new Dyn DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(3): credentials[0] = customer name
//         credentials[1] = username
//         credentials[2] = password
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return dyn.NewDNSProvider()
	case 3:
		config := dyn.NewDefaultConfig()
		config.CustomerName = credentials[0]
		config.UserName = credentials[1]
		config.Password = credentials[2]
		return dyn.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
