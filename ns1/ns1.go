package ns1

import (
	"errors"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/go-acme/lego/v3/providers/dns/ns1"
)

func init() {
	caddytls.RegisterDNSProvider("ns1", NewDNSProvider)
}

// NewDNSProvider returns a new ns1.DNSProvider DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = API key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return ns1.NewDNSProvider()
	case 1:
		config := ns1.NewDefaultConfig()
		config.APIKey = credentials[0]
		return ns1.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
