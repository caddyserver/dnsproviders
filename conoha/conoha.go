// Package conoha adapts the lego ConoHa DNS provider
// for Caddy. Importing this package plugs it in.
package conoha

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/go-acme/lego/providers/dns/conoha"
)

func init() {
	caddytls.RegisterDNSProvider("conoha", NewDNSProvider)
}

// NewDNSProvider returns a new ConoHa DNS challenge provider.
// The credentials are detected automatically; see underlying
// package docs for details:
// https://godoc.org/github.com/go-acme/lego/providers/dns/conoha
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return conoha.NewDNSProvider()
	default:
		return nil, errors.New("invalid credentials length")
	}
}
