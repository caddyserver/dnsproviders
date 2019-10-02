// Package namedotcom adapts the lego Name.com DNS provider
// for Caddy. Importing this package plugs it in.
package namedotcom

import (
	"errors"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/go-acme/lego/v3/providers/dns/namedotcom"
)

func init() {
	caddytls.RegisterDNSProvider("namedotcom", NewDNSProvider)
}

// NewDNSProvider returns a new Name.com DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = username
//         credentials[1] = API token
// len(3): credentials[0] = username
//         credentials[1] = API token
//         credentials[2] = Server
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	config := namedotcom.NewDefaultConfig()

	switch len(credentials) {
	case 0:
		return namedotcom.NewDNSProvider()
	case 3:
		config.Server = credentials[2]
		fallthrough
	case 2:
		config.Username = credentials[0]
		config.APIToken = credentials[1]
		return namedotcom.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
