// Package digitalocean adapts the lego DigitalOcean DNS
// provider for Caddy. Importing this package plugs it in.
package digitalocean

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/go-acme/lego/providers/dns/digitalocean"
)

func init() {
	caddytls.RegisterDNSProvider("digitalocean", NewDNSProvider)
}

// NewDNSProvider returns a new DigitalOcean DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = access token (API key)
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return digitalocean.NewDNSProvider()
	case 1:
		config := digitalocean.NewDefaultConfig()
		config.AuthToken = credentials[0]
		return digitalocean.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
