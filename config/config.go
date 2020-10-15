package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

// EnvManagerConfig represents EnvManager configuration dir env var.
const EnvManagerConfig = "ENV_MANAGER_CONFIG"

var (
	// DefaultEnvManagerHome represent EnvManager home directory.
	DefaultEnvManagerHome = filepath.Join(mustEnvManagerHome(), ".env-manager")
	// EnvManagerConfigFile represents EnvManager config file location.
	EnvManagerConfigFile = filepath.Join(EnvManagerHome(), "config.yml")
	// EnvManagerLogs represents EnvManager log.
	EnvManagerLogs = filepath.Join(os.TempDir(), fmt.Sprintf("env-manager-%s.log", MustEnvManagerUser()))
	// EnvManagerDumpDir represents a directory where EnvManager screen dumps will be persisted.
	EnvManagerDumpDir = filepath.Join(os.TempDir(), fmt.Sprintf("env-manager-screens-%s", MustEnvManagerUser()))
)

type (
	// DeploymentSettings exposes kubeconfig context information.
	DeploymentSettings interface {
		// CurrentDeploymentName returns the name of the current Deployment.
		CurrentDeploymentName() (string, error)

		// CurrentCollectionName returns the name of the current project Collection.
		CurrentCollectionName() (string, error)

		// CurrentProjectName returns the name of the current Project.
		CurrentProjectName() (string, error)

		// CurrentServer returns the name of the current Server.
		CurrentServer() (string, error)

		// DeploymentNames() returns all available Deployment names.
		DeploymentNames() ([]string, error)

		// CollectionNames() returns all available Collection names.
		CollectionNames() ([]string, error)

		// ProjectNames() returns all available project names.
		ProjectNames() ([]string, error)

		// ServerNames returns all available Server names.
		ServerNames() []string
	}

	// Config tracks K9s configuration options.
	Config struct {
		EnvManager *EnvManager
		Settings   DeploymentSettings
	}
)

// EnvManagerHome returns EnvManager configs home directory.
func EnvManagerHome() string {
	if env := os.Getenv(EnvManagerConfig); env != "" {
		return env
	}

	return DefaultEnvManagerHome
}

// NewConfig creates a new default config.
func NewConfig(ks DeploymentSettings) *Config {
	return &Config{Settings: ks}
}

// @TODO Refine the configuration based on cli args.
// func (c *Config) Refine(flags *genericclioptions.ConfigFlags) error {
// 	cfg, err := flags.ToRawKubeConfigLoader().RawConfig()
// 	if err != nil {
// 		return err
// 	}

// 	if isSet(flags.Context) {
// 		c.K9s.CurrentContext = *flags.Context
// 	} else {
// 		c.K9s.CurrentContext = cfg.CurrentContext
// 	}
// 	log.Debug().Msgf("Active Context %q", c.K9s.CurrentContext)
// 	if c.K9s.CurrentContext == "" {
// 		return errors.New("Invalid kubeconfig context detected")
// 	}
// 	context, ok := cfg.Contexts[c.K9s.CurrentContext]
// 	if !ok {
// 		return fmt.Errorf("The specified context %q does not exists in kubeconfig", c.K9s.CurrentContext)
// 	}
// 	c.K9s.CurrentCluster = context.Cluster
// 	if len(context.Namespace) != 0 {
// 		if err := c.SetActiveNamespace(context.Namespace); err != nil {
// 			return err
// 		}
// 	}

// 	if isSet(flags.ClusterName) {
// 		c.K9s.CurrentCluster = *flags.ClusterName
// 	}

// 	if isSet(flags.Namespace) {
// 		if err := c.SetActiveNamespace(*flags.Namespace); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// SetActiveNamespace set the active namespace in the current cluster.
// func (c *Config) SetActiveNamespace(ns string) error {
// 	if c.envManager.ActiveCluster() != nil {
// 		return c.K9s.ActiveCluster().Namespace.SetActive(ns, c.settings)
// 	}
// 	err := errors.New("no active cluster. unable to set active namespace")
// 	log.Error().Err(err).Msg("SetActiveNamespace")

// 	return err
// }

// ActiveView returns the active view in the current cluster.
// func (c *Config) ActiveView() string {
// 	if c.K9s.ActiveCluster() == nil {
// 		return defaultView
// 	}

// 	cmd := c.K9s.ActiveCluster().View.Active
// 	if c.K9s.manualCommand != nil && *c.K9s.manualCommand != "" {
// 		cmd = *c.K9s.manualCommand
// 	}

// 	return cmd
// }

// SetActiveView set the currently cluster active view
// func (c *Config) SetActiveView(view string) {
// 	cl := c.K9s.ActiveCluster()
// 	if cl != nil {
// 		cl.View.Active = view
// 	}
// }

// GetConnection return an api server connection.
// func (c *Config) GetConnection() client.Connection {
// 	return c.client
// }

// SetConnection set an api server connection.
// func (c *Config) SetConnection(conn client.Connection) {
// 	c.client = conn
// }

// Load K9s configuration from file
func (c *Config) Load(path string) error {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	c.EnvManager = NewEnvManager()

	var cfg Config
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return err
	}
	if cfg.EnvManager != nil {
		c.EnvManager = cfg.EnvManager
	}
	// if c.EnvManager.Logger == nil {
	// 	c.EnvManager.Logger = NewLogger()
	// }
	return nil
}

// Save configuration to disk.
func (c *Config) Save() error {
	log.Debug().Msg("[Config] Saving configuration...")
	c.Validate()

	return c.SaveFile(EnvManagerConfigFile)
}

// SaveFile K9s configuration to disk.
func (c *Config) SaveFile(path string) error {
	EnsurePath(path, DefaultDirMod)
	cfg, err := yaml.Marshal(c)
	if err != nil {
		log.Error().Msgf("[Config] Unable to save EnvManager config file: %v", err)
		return err
	}
	return ioutil.WriteFile(path, cfg, 0644)
}

// Validate the configuration.
func (c *Config) Validate() {
	c.EnvManager.Validate( /*c.client, c.settings*/ )
}

// Dump debug...
// func (c *Config) Dump(msg string) {
// 	log.Debug().Msgf("Current Cluster: %s\n", c.K9s.CurrentCluster)
// 	for k, cl := range c.K9s.Clusters {
// 		log.Debug().Msgf("K9s cluster: %s -- %s\n", k, cl.Namespace)
// 	}
// }

// ----------------------------------------------------------------------------
// Helpers...

func isSet(s *string) bool {
	return s != nil && len(*s) > 0
}
