package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	sysinfo := syscall.Sysinfo_t{}

	if err := syscall.Sysinfo(&sysinfo); err != nil {
		fmt.Println(err)
	}

	format := "%s up  1:41,  2 users,  load average: 0,13, 0,32, 0,37"
	fmt.Printf(format, time.Now().Format("15:04:05"))
}
