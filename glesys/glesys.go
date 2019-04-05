// Package glesys adapts the lego GleSYS DNS provider
// for Caddy. Importing this package plugs it in.
package glesys

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/go-acme/lego/providers/dns/glesys"
)

func init() {
	caddytls.RegisterDNSProvider("glesys", NewDNSProvider)
}

// NewDNSProvider returns a new GleSYS DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = API user
//         credentials[1] = API key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return glesys.NewDNSProvider()
	case 2:
		config := glesys.NewDefaultConfig()
		config.APIUser = credentials[0]
		config.APIKey = credentials[1]
		return glesys.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
