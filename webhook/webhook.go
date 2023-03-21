package webhook

// WebHook specifies necessary methods for a webhook service.
type WebHook interface {
	// Start starts the service.
	Start() error

	// Stop terminates the service.
	Stop()

	// Ping checks if the service is alive.
	Ping() error

	// Alert posts an alert message to the WebHook.
	Alert(msg interface{}) error

	// Info posts an info message to the WebHook.
	Info(msg interface{}) error

	// Error posts an error message to the WebHook.
	Error(msg interface{}) error
}
