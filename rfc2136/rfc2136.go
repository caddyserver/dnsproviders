// Package rfc2136 adapts the lego RFC 2136 dynamic update DNS
// provider for Caddy. Importing this package plugs it in.
package rfc2136

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/acme"
	"github.com/xenolf/lego/providers/dns/rfc2136"
)

func init() {
	caddytls.RegisterDNSProvider("rfc2136", NewDNSProvider)
}

// NewDNSProvider returns a new RFC 2136 DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(4): credentials[0] = nameserver
//         credentials[1] = TSIG algorithm
//         credentials[2] = TSIG key
//         credentials[3] = TSIG secret
func NewDNSProvider(credentials ...string) (acme.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return rfc2136.NewDNSProvider()
	case 4:
		return rfc2136.NewDNSProviderCredentials(credentials[0], credentials[1], credentials[2], credentials[3])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
