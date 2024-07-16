package fflags

import (
	"net/http"
	"time"

	"github.com/0x16F/cloud-common/pkg/logger"
	gofeatureflag "github.com/open-feature/go-sdk-contrib/providers/go-feature-flag/pkg"
	of "github.com/open-feature/go-sdk/openfeature"
	"github.com/pkg/errors"
)

// FFlags is a package representing the logic of feature flags.

func NewProvider(log logger.Logger, proxyEndpoint string) (*gofeatureflag.Provider, error) {
	options := gofeatureflag.ProviderOptions{
		Endpoint: proxyEndpoint,
		HTTPClient: &http.Client{
			Timeout: 1 * time.Second,
		},
	}

	provider, err := gofeatureflag.NewProvider(options)
	if err != nil {
		log.Errorf("failed to create feature flag provider: %v", err)

		return nil, errors.Wrap(err, "failed to create feature flag provider")
	}

	return provider, nil
}

func NewClient(log logger.Logger, provider *gofeatureflag.Provider, clientName string) (*of.Client, error) {
	if err := of.SetProvider(provider); err != nil {
		log.Errorf("failed to set provider: %v", err)

		return nil, errors.Wrap(err, "failed to set provider")
	}

	return of.NewClient(clientName), nil
}
