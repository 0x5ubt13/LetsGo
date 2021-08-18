package main

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type PortScanner struct {
	ip   string
	lock *semaphore.Weighted
}
