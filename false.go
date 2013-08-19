package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	help_text = `
	Usage: false [ignored command line arguments]
  	or:  false OPTION
	Exit with a status code indicating failure.

      --help     display this help and exit
      --version  output version information and exit
    `

	help = flag.Bool("help", false, help_text)

	version_text = "false (go-coreutils) 0.1"
	version      = flag.Bool("version", false, version_text)
)

func main() {
	flag.Parse()

	if flag.NFlag() > 1 {
		os.Exit(-1)
	}

	if *help {
		log.Fatal(help_text)
	}

	if *version {
		log.Fatal(version_text)
	}
	os.Exit(-1)
}
