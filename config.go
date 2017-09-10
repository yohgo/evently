package evently

import "github.com/RichardKnop/machinery/v1/config"

// ServerConfig is a wrapper around the machinery config type
type ServerConfig *config.Config

// GetConfiguration retrieves the event server configuration
func GetConfiguration(settings map[string]string) ServerConfig {
	var configuration ServerConfig

	switch settings["source"] {
	case "file":
		configuration = config.NewFromYaml(settings["path"], true, true)
	case "environment":
		configuration = config.NewFromEnvironment(true, true)
	case "manual":
		configuration = &config.Config{
			Broker:        settings["broker"],
			DefaultQueue:  settings["default_queue"],
			ResultBackend: settings["result_backend"],
			AMQP: &config.AMQPConfig{
				Exchange:     settings["amqp_exchange"],
				ExchangeType: settings["amqp_exchange_type"],
				BindingKey:   settings["amqp_binding_key"],
			},
		}
	}

	return configuration
}
