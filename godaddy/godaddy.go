// Package godaddy adapts the lego GoDaddy DNS
// provider for Caddy. Importing this package plugs it in.
package godaddy

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/godaddy"
)

func init() {
	caddytls.RegisterDNSProvider("godaddy", NewDNSProvider)
}

// NewDNSProvider returns a new GoDaddy DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = API key
//         credentials[1] = API secret
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return godaddy.NewDNSProvider()
	case 3:
		return godaddy.NewDNSProviderCredentials(credentials[0], credentials[1])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
