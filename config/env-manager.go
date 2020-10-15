package config

// EnvManager tracks EnvManager configuration options.
type EnvManager struct {
	EnableMouse bool `yaml:"enableMouse"`
	ReadOnly    bool `yaml:"readOnly"`
	// Logger            *Logger             `yaml:"logger"`
	CurrentDeployment  string                  `yaml:"currentDeployment"`
	CurrentCollection  string                  `yaml:"currentCollection"`
	CurrentApplication string                  `yaml:"currentApplication"`
	CurrentServer      string                  `yaml:"currentServer"`
	Servers            map[string]*Server      `yaml:"servers,omitempty"`
	Applications       map[string]*Application `yaml:"applications,omitempty"`
	Deployments        map[string]*Deployment  `yaml:"deployments,omitempty"`
}

// NewEnvManager create a new EnvManager configuration.
func NewEnvManager() *EnvManager {
	return &EnvManager{
		Servers:      make(map[string]*Server),
		Applications: make(map[string]*Application),
		Deployments:  make(map[string]*Deployment),
	}
}

// Validate the current configuration.
func (k *EnvManager) Validate() {
	// k.validateDefaults()
	if k.CurrentDeployment != "" {
		if k.Deployments == nil || k.Deployments[k.CurrentDeployment] == nil {
			k.CurrentDeployment = ""
		}
	}

	if k.CurrentCollection != "" && k.CurrentDeployment != "" {
		if k.Deployments == nil || k.Deployments[k.CurrentDeployment] == nil {
			k.CurrentDeployment = ""
		}
	}

	// if k.Logger == nil {
	// 	k.Logger = NewLogger()
	// } else {
	// 	k.Logger.Validate()
	// }
}
