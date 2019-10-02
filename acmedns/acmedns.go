// Package acmedns adapts the lego exec DNS
// provider for Caddy. Importing this package plugs it in.
package acmedns

import (
	"errors"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/go-acme/lego/v3/providers/dns/acmedns"
)

func init() {
	caddytls.RegisterDNSProvider("acmedns", NewDNSProvider)
}

// NewDNSProvider returns a new acmedns DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return acmedns.NewDNSProvider()
	default:
		return nil, errors.New("invalid credentials length")
	}
}
