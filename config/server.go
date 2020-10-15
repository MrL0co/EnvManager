package config

// Server tracks server configuration
type Server struct {
	IP         string `yaml:"ip"`
	User       string `yaml:"user"`
	SSHKeyFile string `yaml:"sshKeyFile"`
	password   string `yaml:"password"`
}

// NewServer creates a new Server configuration.
func NewServer() *Server {
	return &Server{
		IP:         "",
		User:       "",
		SSHKeyFile: nil,
		password:   nil,
	}
}
