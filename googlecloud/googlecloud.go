// Package googlecloud adapts the lego Google Cloud DNS
// provider for Caddy. Importing this package plugs it in.
package googlecloud

import (
	"errors"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/go-acme/lego/providers/dns/gcloud"
)

func init() {
	caddytls.RegisterDNSProvider("googlecloud", NewDNSProvider)
}

// NewDNSProvider returns a new Google Cloud DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = project
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return gcloud.NewDNSProvider()
	case 1:
		config := gcloud.NewDefaultConfig()
		config.Project = credentials[0]
		return gcloud.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
