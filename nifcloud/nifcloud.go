// Package nifcloud adapts the lego NIFCLOUD DNS
// provider for Caddy. Importing this package plugs it in.
package nifcloud

import (
	"errors"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/go-acme/lego/v3/providers/dns/nifcloud"
)

func init() {
	caddytls.RegisterDNSProvider("nifcloud", NewDNSProvider)
}

// NewDNSProvider returns a new NIFCLOUD DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment (https://godoc.org/github.com/go-acme/lego/providers/dns/nifcloud)
// len(3): credentials[0] = Base URL
//         credentials[1] = ACCESS KEY ID
//         credentials[2] = SECRET ACCESS KEY
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return nifcloud.NewDNSProvider()
	case 3:
		config := nifcloud.NewDefaultConfig()
		config.BaseURL = credentials[0]
		config.AccessKey = credentials[1]
		config.SecretKey = credentials[2]
		return nifcloud.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
