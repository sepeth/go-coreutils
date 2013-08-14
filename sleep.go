package main

import (
	"flag"
	"os"
	"strconv"
	"strings"
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

	i := flag.Arg(0)
	sec, _ := strconv.Atoi(strings.TrimRight(i, "s"))
	min, _ := strconv.Atoi(strings.TrimRight(i, "m"))
	day, _ := strconv.Atoi(strings.TrimRight(i, "d"))

	if strings.Contains(i, "s") {
		time.Sleep(time.Duration(sec) * time.Second)
		return
	}

	if strings.Contains(i, "m") {
		time.Sleep(time.Duration(min) * time.Minute)
		return
	}

	if strings.Contains(i, "d") {
		time.Sleep(time.Duration(day) * time.Second * 86400)
		return
	}

	time.Sleep(time.Duration(i) * time.Second)

}
