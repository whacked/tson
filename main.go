package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"syscall"

	"github.com/mitchellh/go-homedir"
	"github.com/whacked/tson/gui"
	"golang.org/x/term"
)

var (
	enableLog = flag.Bool("log", false, "enable log")
	url       = flag.String("url", "", "get json from url")
)

func printError(err error) int {
	fmt.Fprintln(os.Stderr, err)
	return 1
}

func init() {
	flag.Parse()
	if *enableLog {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		logWriter, err := os.OpenFile(filepath.Join(home, "tson.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			os.Exit(printError(err))
		}

		log.SetOutput(logWriter)
		log.SetFlags(log.Lshortfile)
	} else {
		log.SetOutput(io.Discard)
	}

}

func run() int {
	var i interface{}
	if *url != "" {
		resp, err := http.Get(*url)
		if err != nil {
			return printError(err)
		}
		defer resp.Body.Close()
		i, err = gui.UnMarshalJSON(resp.Body)
		if err != nil {
			return printError(err)
		}
	} else {
		if !term.IsTerminal(int(os.Stdin.Fd())) {
			var err error
			i, err = gui.UnMarshalJSON(os.Stdin)
			if err != nil {
				return printError(err)
			}

			// set tview tty to stdin
			os.Stdin = os.NewFile(uintptr(syscall.Stderr), "/dev/tty")
		}
	}

	if i == nil {
		return printError(gui.ErrEmptyJSON)
	}

	if err := gui.New().Run(i); err != nil {
		log.Println(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
