// Package dyn adapts the lego Dyn DNS provider
// for Caddy. Importing this package plugs it in.
package dyn

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/acme"
	"github.com/xenolf/lego/providers/dns/dyn"
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
func NewDNSProvider(credentials ...string) (acme.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return dyn.NewDNSProvider()
	case 3:
		return dyn.NewDNSProviderCredentials(credentials[0], credentials[1], credentials[2])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
