package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	groups = make(map[string]string)

	help_text = `
	Usage: groups [OPTION]... [USERNAME]...
	Print group memberships for each USERNAME or, if no USERNAME is specified, for
	the current process (which may differ if the groups database has changed).
      --help     display this help and exit
      --version  output version information and exit
	`
	help = flag.Bool("help", false, help_text)

	version_text = "groups (go-coreutils) 0.1"
	version      = flag.Bool("version", false, version_text)
)

func print_groups(user string, method int) {
	ret := user

	for g, u := range groups {
		if u == user {
			ret = ret + " " + g
		}
	}

	if method == 0 {
		fmt.Println(ret)

	} else {
		fmt.Println(user + " : " + ret)
	}

	return
}

func main() {
	flag.Parse()

	if *help {
		fmt.FPrintln(os.Stderr, help_text)
		return
	}

	if *version {
		fmt.FPrintln(os.Stderr, version_text)
		return
	}

	file, err := os.Open("/etc/group")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	file_stats, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, file_stats.Size())
	file.Read(data)

	data_n := strings.Split(string(data), "\n")
	for _, v := range data_n {
		x := strings.Split(v, ":")
		if len(x) >= 3 && len(x[3]) > 0 {
			groups[x[0]] = x[3]
		}
	}

	if len(flag.Args()) == 1 {
		user := flag.Arg(0)
		print_groups(user, 1)
		return
	}

	if len(flag.Args()) == 0 {
		user := os.Getenv("USER")
		print_groups(user, 0)
		return
	}
}
