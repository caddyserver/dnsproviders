// Package linode adapts the lego Linode DNS
// provider for Caddy. Importing this package plugs it in.
package linode

import (
	"errors"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/go-acme/lego/providers/dns/linode"
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
		config := linode.NewDefaultConfig()
		config.APIKey = credentials[0]
		return linode.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
