// Package azure adapts the lego azure DNS
// provider for Caddy. Importing this package plugs it in.
package azure

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/azure"
)

func init() {
	caddytls.RegisterDNSProvider("azure", NewDNSProvider)
}

// NewDNSProvider returns a new azure DNS challenge provider.
// The credentials are detected automatically; see underlying
// package docs for details:
// https://godoc.org/github.com/xenolf/lego/providers/dns/azure
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return azure.NewDNSProvider()
	default:
		return nil, errors.New("invalid credentials length")
	}
}
