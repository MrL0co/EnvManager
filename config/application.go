package config

// Application tracks Application configuration
type Application struct {
	EnvVars []string
}

// NewApplication creates a new Application configuration.
func NewApplication() *Application {
	return &Application{
		EnvVars: []string{},
	}
}
