package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"path/filepath"

	"git.sr.ht/~otl/pushover"
)

const usage string = "usage: pover [-d] [-f file]"

var debug *bool
var configflag *string

func init() {
	debug = flag.Bool("d", false, "debug")
	configflag = flag.String("f", "", "path to configuration file")
	flag.Parse()
}

func main() {
	var configpath string
	configpath = *configflag
	if *configflag == "" {
		s, err := os.UserConfigDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		configpath = filepath.Join(s, "pover")
	}
	config, err := configFromFile(configpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load configuration: %v\n", err)
		os.Exit(1)
	}

	lr := io.LimitReader(os.Stdin, pushover.MaxMsgLength)
	b, err := io.ReadAll(lr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(b) == pushover.MaxMsgLength {
		fmt.Fprintf(os.Stderr, "max message length (%d) reached\n", pushover.MaxMsgLength)
	}
	if *debug {
		fmt.Fprint(os.Stderr, string(b))
	}

	msg := pushover.Message{
		User: config.user,
		Token: config.token,
		Message: string(b),
	}
	if err := pushover.Push(msg); err != nil {
		fmt.Fprintf(os.Stderr, "push message: %v\n", err)
		os.Exit(1)
	}
}
