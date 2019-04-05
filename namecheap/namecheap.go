// Package namecheap adapts the lego NameCheap DNS provider
// for Caddy. Importing this package plugs it in.
package namecheap

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/go-acme/lego/providers/dns/namecheap"
)

func init() {
	caddytls.RegisterDNSProvider("namecheap", NewDNSProvider)
}

// NewDNSProvider returns a new NameCheap DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = API user
//         credentials[1] = API key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return namecheap.NewDNSProvider()
	case 2:
		config := namecheap.NewDefaultConfig()
		config.APIUser = credentials[0]
		config.APIKey = credentials[1]
		return namecheap.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
