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
		config := fastdns.NewDefaultConfig()
		config.Config.Host = credentials[0]
		config.Config.ClientToken = credentials[1]
		config.Config.ClientSecret = credentials[2]
		config.Config.AccessToken = credentials[3]
		return fastdns.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
