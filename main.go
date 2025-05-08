package main

import (
	"cast/cmd"

	_ "cast/cmd/experimental"
	_ "cast/cmd/fixes"
	_ "cast/cmd/umbra"
)

// "cast/cmd"
// _ "cast/cmd/devtools"
// _ "cast/cmd/horizon"

// _ "cast/cmd/nova"
func main() {
	// config.LoadConfig()

	cmd.Execute()
}
