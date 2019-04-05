package stackpath

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/go-acme/lego/providers/dns/stackpath"
)

func TestNewDNSProvider(t *testing.T) {
	os.Setenv("STACKPATH_CLIENT_ID", "client-id-env")
	os.Setenv("STACKPATh_CLIENT_SECRET", "client-idsecret-env")
	os.Setenv("STACKPATH_STACK_ID", "stack-id-env")

	tests := map[string]struct {
		credentials []string
		want        func() (*stackpath.DNSProvider, error)
		expectedErr error
	}{
		"no_creds_given": {
			credentials: []string{},
			want: func() (*stackpath.DNSProvider, error) {
				return stackpath.NewDNSProvider()
			},
			expectedErr: nil,
		},
		"creds_given": {
			credentials: []string{
				"client-id-given",
				"client-secret-given",
				"stack-id-given",
			},
			want: func() (*stackpath.DNSProvider, error) {
				config := stackpath.NewDefaultConfig()
				config.ClientID = "client-id-given"
				config.ClientSecret = "client-secret-given"
				config.StackID = "stack-id-given"
				return stackpath.NewDNSProviderConfig(config)
			},
			expectedErr: nil,
		},
		"bad_cred_count": {
			credentials: []string{
				"client-id-given",
				"client-secret-given",
			},
			want: func() (*stackpath.DNSProvider, error) {
				return nil, nil
			},
			expectedErr: errors.New("invalid credentials length"),
		},
	}
	for ttName, tt := range tests {
		t.Run(ttName, func(t *testing.T) {
			expectedProvider, err := tt.want()
			if err != nil {
				t.Error(err)
			}

			got, err := NewDNSProvider(tt.credentials...)
			if err == nil && tt.expectedErr != nil {
				t.Errorf("did not get expected error: %v", tt.expectedErr)
				return
			}

			if err != nil && tt.expectedErr == nil {
				t.Errorf("unexpected err: %v", err)
				return
			}

			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("did not get error that was expected. got: %v, want: %v", err, tt.expectedErr)
				return
			}

			if err == nil && !reflect.DeepEqual(got, expectedProvider) {
				t.Errorf("unexpected result. got: %v, want: %v", got, expectedProvider)
			}
		})
	}
}
