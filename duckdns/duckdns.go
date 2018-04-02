// Package duckdns adapts the lego duckdns DNS
// provider for Caddy. Importing this package plugs it in.
package duckdns

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/duckdns"
)

func init() {
	caddytls.RegisterDNSProvider("duckdns", NewDNSProvider)
}

// NewDNSProvider returns a new duckdns DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = duckdns token
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return duckdns.NewDNSProvider()
	case 1:
		return duckdns.NewDNSProviderCredentials(credentials[0])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
