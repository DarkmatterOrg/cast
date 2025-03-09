package main

import (
	"cast/cmd"
	_ "cast/cmd/devtools"
	_ "cast/cmd/horizon"
	_ "cast/cmd/nova"
)

func main() {
	cmd.Execute()
}
