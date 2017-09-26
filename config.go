package evently

import (
	"strings"

	"github.com/RichardKnop/machinery/v1/config"
)

// Configuration is a wrapper around the machinery config type
type Configuration struct {
	ServerConfig   *config.Config
	UpstreamQueues []string
}

// GetConfiguration retrieves the event server configuration
func GetConfiguration(settings map[string]string) Configuration {
	return Configuration{
		ServerConfig: &config.Config{
			Broker:        settings["broker"],
			DefaultQueue:  settings["downstream_queue"],
			ResultBackend: settings["result_backend"],
			AMQP: &config.AMQPConfig{
				Exchange:     settings["amqp_exchange"],
				ExchangeType: settings["amqp_exchange_type"],
				BindingKey:   settings["amqp_binding_key"],
			},
		},
		UpstreamQueues: strings.Split(settings["upstream_queues"], ";"),
	}
}
