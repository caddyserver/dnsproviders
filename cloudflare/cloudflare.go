// Package cloudflare adapts the lego Cloudflare DNS
// provider for Caddy. Importing this package plugs it in.
package cloudflare

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/cloudflare"
)

func init() {
	caddytls.RegisterDNSProvider("cloudflare", NewDNSProvider)
}

// NewDNSProvider returns a new Cloudflare DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = Email address
//         credentials[1] = API key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return cloudflare.NewDNSProvider()
	case 2:
		return cloudflare.NewDNSProviderCredentials(credentials[0], credentials[1])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
