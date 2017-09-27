package evently

import (
	machinery "github.com/RichardKnop/machinery/v1"
)

// EventListener is an application event listener.
type EventListener struct {
	Server   *machinery.Server
	Listener *machinery.Worker
	Errors   []error
}

// NewEventListener creates a new application event listener.
func NewEventListener(name string, settings map[string]string, handlers map[string]interface{}) *EventListener {
	listener := &EventListener{}
	var err error

	// Attempting to start the event server
	listener.Server, err = machinery.NewServer(GetConfiguration(settings).ServerConfig)
	listener.processError(err)

	if listener.IsOK() {
		// Creating the application event listener
		listener.Listener = listener.Server.NewWorker(name, 0)

		// Registering the event tasks with the event server
		err = listener.Server.RegisterTasks(handlers)
		listener.processError(err)
	}

	return listener
}

// Eavesdrop tells the event listener to start listening for certain events
func (listener *EventListener) Eavesdrop() *EventListener {
	go listener.Listener.Launch()
	return listener
}

// processError processes an event listener error
func (listener *EventListener) processError(err error) {
	if err != nil {
		listener.Errors = append(listener.Errors, err)
	}
}

// IsOK determins the ok status of the event listener
func (listener *EventListener) IsOK() bool {
	return (len(listener.Errors) == 0)
}
