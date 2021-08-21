package main

import (
	"fmt"
	"golang.org/x/sync/semaphore"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type PortScanner struct {
	// The host to scan
	ip string

	// The threshold to limit the goroutines
	lock *semaphore.Weighted
}

func Ulimit() int64 {
	// Call the built-in command `ulimit` to help lock
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func ScanPort(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		// Handling socket error
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
			//
		} else {
			fmt.Println(port, "closed")
		}

		return
	}

	conn.Close()
	fmt.Println(port, "open")
}
