package main

import (
	"flag"
	"os"
	"strconv"
	"time"
)

func usage() {
	os.Stderr.WriteString("Usage: sleep seconds")
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		usage()
		os.Exit(1)
	}
	i, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		usage()
		os.Exit(1)
	}
	time.Sleep(time.Duration(i) * time.Second)
}
