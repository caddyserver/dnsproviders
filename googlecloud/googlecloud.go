// Package googlecloud adapts the lego Google Cloud DNS
// provider for Caddy. Importing this package plugs it in.
package googlecloud

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/acme"
	"github.com/xenolf/lego/providers/dns/googlecloud"
)

func init() {
	caddytls.RegisterDNSProvider("googlecloud", NewDNSProvider)
}

// NewDNSProvider returns a new Google Cloud DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = project
func NewDNSProvider(credentials ...string) (acme.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return googlecloud.NewDNSProvider()
	case 1:
		return googlecloud.NewDNSProviderCredentials(credentials[0])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
