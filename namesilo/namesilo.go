// Package namesilo adapts the lego namesilo DNS provider
// for Caddy. Importing this package plugs it in.
package namesilo

import (
	"errors"
	"strconv"
	"time"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/go-acme/lego/v3/providers/dns/namesilo"
)

func init() {
	caddytls.RegisterDNSProvider("namesilo", NewDNSProvider)
}

// NewDNSProvider returns a new namesilo DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(4): credentials[0] = API key
//         credentials[1] = TTL
//         credentials[2] = Propagation timeout
//         credentials[3] = Polling interval
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return namesilo.NewDNSProvider()
	case 4:
		config := namesilo.NewDefaultConfig()
		config.APIKey = credentials[0]
		if ttl, err := strconv.ParseInt(credentials[1], 10, 32); err != nil {
			return nil, err
		} else {
			config.TTL = int(ttl)
		}
		if timeout, err := time.ParseDuration(credentials[2]); err != nil {
			return nil, err
		} else {
			config.PropagationTimeout = timeout
		}
		if interval, err := time.ParseDuration(credentials[3]); err != nil {
			return nil, err
		} else {
			config.PollingInterval = interval
		}
		return namesilo.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
