package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
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
			// Handling closed port
		} else {
			fmt.Println(port, "closed")
		}

		return
	}

	conn.Close()
	fmt.Println(port, "open")
}

func (ps *PortScanner) Start(f, l int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := f; port <= l; port++ {
		wg.Add(1)
		ps.lock.Acquire(context.TODO(), 1)

		go func(port int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			ScanPort(ps.ip, port, timeout)
		}(port)
	}
}

func main() {
	ps := &PortScanner{
		ip:   "127.0.0.1",
		lock: semaphore.NewWeighted(Ulimit()),
	}

	ps.Start(1, 65535, 500*time.Millisecond)
}
