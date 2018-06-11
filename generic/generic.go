// Package generic adapts the lego exec DNS
// provider for Caddy. Importing this package plugs it in.
package generic

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/exec"
)

func init() {
	caddytls.RegisterDNSProvider("generic", NewDNSProvider)
}

// NewDNSProvider returns a new exec DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use program to run from environment variable EXEC_PATH
// len(1): credentials[0] = program to run
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return exec.NewDNSProvider()
	case 1:
		return exec.NewDNSProviderProgram(credentials[0])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
