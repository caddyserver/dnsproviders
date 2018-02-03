// Package cloudxns adapts the lego CloudXNS DNS
// provider for Caddy. Importing this package plugs it in.
package cloudxns

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/cloudxns"
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
		return cloudxns.NewDNSProviderCredentials(credentials[0], credentials[1])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
