// Package namecheap adapts the lego NameCheap DNS provider
// for Caddy. Importing this package plugs it in.
package namecheap

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/acme"
	"github.com/xenolf/lego/providers/dns/namecheap"
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
func NewDNSProvider(credentials ...string) (acme.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return namecheap.NewDNSProvider()
	case 2:
		return namecheap.NewDNSProviderCredentials(credentials[0], credentials[1])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
