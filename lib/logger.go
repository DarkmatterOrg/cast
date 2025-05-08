package lib

import (
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

const SuccessLevel = log.InfoLevel + 1
const ImportantWarnLevel = log.WarnLevel + 1

type NewLog struct {
	*log.Logger
}

func (l *NewLog) Success(msg string, args ...any) {
	l.Log(SuccessLevel, msg, args...)
}

func (l *NewLog) ImportantWarn(msg string, args ...any) {
	style := lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("192")).
		Foreground(lipgloss.Color("0"))

	l.Log(ImportantWarnLevel, style.Render(msg), args...)
}

func NewLogger() *NewLog {
	l := new(NewLog)
	logger := log.NewWithOptions(os.Stderr, log.Options{
			ReportTimestamp: true,
			TimeFormat: time.TimeOnly,
			Level: log.DebugLevel,
		})

	styles := log.DefaultStyles()

	styles.Levels[SuccessLevel] = lipgloss.NewStyle().
		SetString("SUCCESS").
		Bold(true).
		MaxWidth(4).
		Foreground(lipgloss.Color("85"))

	styles.Levels[ImportantWarnLevel] = lipgloss.NewStyle().
		SetString("WARN").
		Bold(true).
		MaxWidth(4).
		Foreground(lipgloss.Color("192"))

	logger.SetStyles(styles)

	l.Logger = logger

	return l
}

var Logger = NewLogger()