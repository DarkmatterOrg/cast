package util

import (
	"os"
)

func CheckEnableDisableFlag(enableFlag bool, disableFlag bool) {
	if enableFlag && disableFlag {
		if Config.Insult {
			Logger.Warn("You can't use both --enable and --disable at the same time you fucking moron!")
		} else {
			Logger.Warn("You can't use both --enable and --disable at the same time.")
		}
		os.Exit(0)
	}
}