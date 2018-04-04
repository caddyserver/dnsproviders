// Package fastdsn adapts the lego FastDNS DNS provider
// for Caddy. Importing this package plugs it in.
package fastdns

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/fastdns"
)

func init() {
	caddytls.RegisterDNSProvider("fastdns", NewDNSProvider)
}

// NewDNSProvider returns a new FastDNS DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(4): credentials[0] = Akamai Host
//         credentials[1] = Client Token
//         credentials[2] = Client Secret
//         credentials[3] = Access Token
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return fastdns.NewDNSProvider()
	case 4:
		return fastdns.NewDNSProviderClient(credentials[0], credentials[1], credentials[2], credentials[3])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
