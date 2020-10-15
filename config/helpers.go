package config

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

const (
	// DefaultDirMod default unix perms for EnvManager directory.
	DefaultDirMod os.FileMode = 0755
	// DefaultFileMod default unix perms for EnvManager files.
	DefaultFileMod os.FileMode = 0600
)

// InList check if string is in a collection of strings.
func InList(ll []string, n string) bool {
	for _, l := range ll {
		if l == n {
			return true
		}
	}
	return false
}

func mustEnvManagerHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal().Err(err).Msg("Die on retrieving user home")
	}
	return usr.HomeDir
}

// MustEnvManagerUser establishes current user identity or fail.
func MustEnvManagerUser() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal().Err(err).Msg("Die on retrieving user info")
	}
	return usr.Username
}

// EnsurePath ensures a directory exist from the given path.
func EnsurePath(path string, mod os.FileMode) {
	dir := filepath.Dir(path)
	EnsureFullPath(dir, mod)
}

// EnsureFullPath ensures a directory exist from the given path.
func EnsureFullPath(path string, mod os.FileMode) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, mod); err != nil {
			log.Fatal().Msgf("Unable to create dir %q %v", path, err)
		}
	}
}
