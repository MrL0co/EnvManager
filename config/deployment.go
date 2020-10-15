package config

// Deployment tracks Deployment configuration
type Deployment struct {
	Servers     map[string]*Server     `yaml:"servers"`
	Folder      string                 `yaml:"folder"`
	Collections map[string]*Collection `yaml:"collections"`
}

// NewDeployment creates a new Deployment configuration.
func NewDeployment() *Deployment {
	return &Deployment{
		Servers:     map[string]*Server{},
		Folder:      "",
		Collections: map[string]*Collection{},
	}
}
