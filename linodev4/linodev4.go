// Package linodev4 adapts the lego LinodeV4 DNS
// provider for Caddy. Importing this package plugs it in.
package linodev4

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/linodev4"
)

func init() {
	caddytls.RegisterDNSProvider("linodev4", NewDNSProvider)
}

// NewDNSProvider returns a new LinodeV4 DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = access token (API token)
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return linodev4.NewDNSProvider()
	case 1:
		config := linodev4.NewDefaultConfig()
		config.Token = credentials[0]
		return linodev4.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
