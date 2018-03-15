// Package namedotcom adapts the lego Name.com DNS provider
// for Caddy. Importing this package plugs it in.
package namedotcom

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/namedotcom"
)

func init() {
	caddytls.RegisterDNSProvider("namedotcom", NewDNSProvider)
}

// NewDNSProvider returns a new Name.com DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = username
//         credentials[1] = API token
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return namedotcom.NewDNSProvider()
	case 2:
		return namedotcom.NewDNSProviderCredentials(credentials[0], credentials[1])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
