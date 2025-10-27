package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/charmbracelet/lipgloss"
)

const (
	DEV  = "dev"
	PROD = "prod"
)

var env = DEV

var (
	INFO    = []string{"#fa9ebc", "INFO  :"}
	WARNING = []string{"#C7B965", "WARN  :"}
	ERROR   = []string{"#911746", "ERROR :"}
	PANIK   = []string{"#751439", "PANIK :"}
	CRASH   = []string{"#B57931", "CRASH :"}
	HANDLE  = []string{"#ED95A4", "HANDLE:"}
)

func log(logType []string, format string, a ...any) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color(logType[0]))
	msg := style.Render(logType[1])
	formattedMsg := fmt.Sprintf(format, a...)
	fmt.Printf("%s %s", msg, formattedMsg)
}

func slog(logType []string, format string, a ...any) string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color(logType[0]))
	msg := style.Render(logType[1])
	formattedMsg := fmt.Sprintf(format, a...)
	return fmt.Sprintf("%s %s", msg, formattedMsg)
}

func logln(logType []string, format string, a ...any) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color(logType[0]))

	msg := style.Render(logType[1])
	formattedMsg := fmt.Sprintf(format, a...)
	fmt.Printf("%s %s\n", msg, formattedMsg)
}

func styleslog(logType []string, format string, a ...any) string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color(logType[0])).
		PaddingLeft(1)
	msg := style.Render(logType[1])
	formattedMsg := fmt.Sprintf(format, a...)
	return fmt.Sprintf("%s %s\n", msg, formattedMsg)
}

func logError(format string, a ...any) {
	logln(ERROR, format, a...)
}

func prettyLogErr(err error, msg string) {
	if env == DEV {
		if err != nil {
			logln(ERROR, "%s", err.Error())
		}
	} else {
		logln(ERROR, "%s", msg)
	}
}

func logErr(err error) {
	if err != nil {
		logln(ERROR, "%s", err.Error())
	}
}

func panik(format string, a ...any) {
	if env != DEV {
		log(WARNING, "consider removing this intentional crash/panik")
	}
	panic(slog(PANIK, format, a...))
	// pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	// os.Exit(1)
}

func panikif(cond bool, format string, a ...any) {
	if !cond {
		return
	}
	if env != DEV {
		log(WARNING, "consider removing this intentional crash/panik")
	}
	log(PANIK, format, a...)
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	os.Exit(0)
}

func userPanikErr(err error) {
	logln(PANIK, "%s", err.Error())
	// TODO: reverse order stack trace please
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	os.Exit(0)
}

func crash(format string, a ...any) {
	logln(CRASH, format, a...)
	os.Exit(0)
}

func exitOnError(format string, a ...any) {
	logln(ERROR, format, a...)
	os.Exit(1)
}

func exit() {
	os.Exit(0)
}

func crashErr(err error) {
	if env != "dev" {
		log(WARNING, "consider removing this intentional crash/panik")
	}
	logln(CRASH, "%s", err.Error())
	os.Exit(0)
}

func crashIfErr(err error) {
	if err != nil {
		if env != "dev" {
			log(WARNING, "consider removing this intentional crash/panik")
		}
		logln(CRASH, "%s", err.Error())
		os.Exit(0)
	}
}

func panikIfErr(err error) {
	if err != nil {
		if env != "dev" {
			log(WARNING, "consider removing this intentional crash/panik")
		}
		panik("%s", err.Error())
	}
}

func logInfo(format string, a ...any) {
	log(INFO, format, a...)
}

func use(a any) any {
	return a
}

func assertThisIsTrue(cond bool, format string, a ...any) {
	if !cond {
		if env != "dev" {
			log(WARNING, "consider removing this intentional crash/panik")
		}
		panik(format, a...)
	}
}
func logInfoln(format string, a ...any) {
	logln(INFO, format, a...)
}
