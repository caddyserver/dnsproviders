// Package dnsimple adapts the lego DNSimple DNS provider
// for Caddy. Importing this package plugs it in.
package dnsimple

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/dnsimple"
)

func init() {
	caddytls.RegisterDNSProvider("dnsimple", NewDNSProvider)
}

// NewDNSProvider returns a new DNSimple DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = email
//         credentials[1] = API key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return dnsimple.NewDNSProvider()
	case 2:
		return dnsimple.NewDNSProviderCredentials(credentials[0], credentials[1])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
