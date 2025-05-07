package components

import (
	"cast/config"
	"cast/lib"
	"os"
)

func CheckEnableDisableFlag(enableFlag bool, disableFlag bool) {
	if enableFlag && disableFlag {
		if config.Config.Insult {
			lib.Logger.Warn("You can't use both --enable and --disable at the same time you fucking moron!")
		} else {
			lib.Logger.Warn("You can't use both --enable and --disable at the same time.")
		}
		os.Exit(0)
	}
}