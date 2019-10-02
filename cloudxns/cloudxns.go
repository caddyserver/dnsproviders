// Package cloudxns adapts the lego CloudXNS DNS
// provider for Caddy. Importing this package plugs it in.
package cloudxns

import (
	"errors"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/go-acme/lego/v3/providers/dns/cloudxns"
)

func init() {
	caddytls.RegisterDNSProvider("cloudxns", NewDNSProvider)
}

// NewDNSProvider returns a new CloudXNS DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = API key
//         credentials[1] = Secret key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return cloudxns.NewDNSProvider()
	case 2:
		config := cloudxns.NewDefaultConfig()
		config.APIKey = credentials[0]
		config.SecretKey = credentials[1]
		return cloudxns.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
