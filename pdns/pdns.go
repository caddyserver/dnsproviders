// Package pdns adapts the lego PowerDNS
// provider for Caddy. Importing this package plugs it in.
package pdns

import (
	"errors"
	"net/url"

	"github.com/mholt/caddy/caddytls"
	"github.com/go-acme/lego/providers/dns/pdns"
)

func init() {
	caddytls.RegisterDNSProvider("powerdns", NewDNSProvider)
}

// NewDNSProvider returns a new PowerDNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = pdns API URL, credentials[1] = pdns API key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return pdns.NewDNSProvider()
	case 2:
		var err error
		config := pdns.NewDefaultConfig()
		config.Host, err = url.Parse(credentials[0])
		if err != nil {
			return nil, errors.New("invalid URL format")
		}
		config.APIKey = credentials[1]
		return pdns.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
