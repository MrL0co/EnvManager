package main

import "fmt"

func init() {
	config.EnsurePath(config.K9sLogs, config.DefaultDirMod)
}

func main() {
	fmt.Println("Hello, world.")
}
