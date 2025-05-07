package lib

import (
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var Logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	TimeFormat: time.TimeOnly,
})

func ImportantWarn(msg string) {
	styles := log.DefaultStyles()

	styles.Message = lipgloss.NewStyle().
	Bold(true).
	Padding(0, 1, 0, 1).
	Background(lipgloss.Color("192")).
	Foreground(lipgloss.Color("0"))


	Logger.SetStyles(styles)

	Logger.Warn(msg)
	
	styles.Message = lipgloss.NewStyle().
	UnsetBold().
	UnsetBackground()
}