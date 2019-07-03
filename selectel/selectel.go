// Package selectel adapts the lego Selectel DNS
// provider for Caddy. Importing this package plugs it in.
package selectel

import (
	"errors"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/go-acme/lego/providers/dns/selectel"
)

func init() {
	caddytls.RegisterDNSProvider("selectel", NewDNSProvider)
}

// NewDNSProvider returns a new Selectel DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment (https://godoc.org/github.com/go-acme/lego/providers/dns/selectel)
// len(1): credentials[0] = Token
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return selectel.NewDNSProvider()
	case 1:
		config := selectel.NewDefaultConfig()
		config.Token = credentials[0]
		return selectel.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
