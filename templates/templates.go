package templates

const (
	Main = `
	package main
	import "fmt"

	func main()  {
		fmt.Println("hello world")
	}
	`

	Makefile = `compile:
	go build -o ./bin/%[1]v -gcflags='all=-N -l' ./cmd/%[1]v   2> ./errors.err
	
c:
	go build -o ./bin/%[1]v -gcflags='all=-N -l' ./cmd/%[1]v
	
run:
	env --chdir=./test ../bin/%[1]v steins_gate lampda

compile_test:
	go test -v ./cmd/%[1]v 2> ./errors.err

test:
	go test -v ./cmd/%[1]v
	
test_f:
	go test -v -run $(FN) ./cmd/%[1]v

debug:debug
	env --chdir=./cmd/%[1]v gdlv debug
	
debugt:debugt
	env --chdir=./cmd/%[1]v gdlv test
	
debug_f:
	env --chdir=./cmd/%[1]v gdlv test -run $(FN)`
	Logger = `
package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/charmbracelet/lipgloss"
)

const (
	DEV = "dev"
)

var env = "dev"

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

func logErr(err error) {
	logln(ERROR, "%s", err.Error())
}

func panik(format string, a ...any) {
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

func logInfo(format string, a ...any) {
	log(INFO, format, a...)
}

func logInfoln(format string, a ...any) {
	logln(INFO, format, a...)
}`
)
