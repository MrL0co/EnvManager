package config

// Collection tracks Collection configuration
type Collection struct {
	Collections  map[string]*Collection  `yaml:"collections"`
	Applications map[string]*Application `yaml:"applications"`
}

// NewCollection creates a new Collection configuration.
func NewCollection() *Collection {
	return &Collection{
		Collections:  map[string]*Collection{},
		Applications: map[string]*Application{},
	}
}
