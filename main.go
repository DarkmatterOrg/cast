package main

import (
	"cast/cmd"
	_ "cast/cmd/devtools"
	_ "cast/cmd/fixes"
	_ "cast/cmd/horizon"
	_ "cast/cmd/nova"
	_ "cast/cmd/umbra"
)

func main() {
	cmd.Execute()
}
