// Package vultr adapts the lego Vultr DNS provider
// for Caddy. Importing this package plugs it in.
package vultr

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/vultr"
)

func init() {
	caddytls.RegisterDNSProvider("vultr", NewDNSProvider)
}

// NewDNSProvider returns a new Vultr DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = API key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return vultr.NewDNSProvider()
	case 1:
		return vultr.NewDNSProviderCredentials(credentials[0])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
