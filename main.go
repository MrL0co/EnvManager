package main

import (
	"os"

	"github.com/MrL0co/env-manager/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	config.EnsurePath(config.EnvManagerLogs, config.DefaultDirMod)
}

func main() {
	mod := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	file, err := os.OpenFile(config.EnvManagerLogs, mod, config.DefaultFileMod)
	if err != nil {
		panic(err)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: file})

}
