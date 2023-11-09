package nanoleafreceiver

import (
	"fmt"

	"github.com/Elfo404/home-otel-collector/pkg/receiver/nanoleafreceiver/internal/metadata"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

// Config defines configuration for nanoleafreceiver receiver.
type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	metadata.MetricsBuilderConfig           `mapstructure:",squash"`
	Endpoint                                string `mapstructure:"endpoint"`
	Token                                   string `mapstructure:"token"`
}

// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error {
	if cfg.Endpoint == "" {
		return fmt.Errorf("endpoint must be defined")
	}
	if cfg.Token == "" {
		return fmt.Errorf("token must be defined")
	}

	return nil
}
