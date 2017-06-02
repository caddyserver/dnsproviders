// Package linode adapts the lego Linode DNS
// provider for Caddy. Importing this package plugs it in.
package linode

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/linode"
)

func init() {
	caddytls.RegisterDNSProvider("linode", NewDNSProvider)
}

// NewDNSProvider returns a new Linode DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = access token (API key)
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return linode.NewDNSProvider()
	case 1:
		return linode.NewDNSProviderCredentials(credentials[0])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
