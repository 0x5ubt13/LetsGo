package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	pass1 := completeLevel(0, getBandit0Password(), "cat readme")
	fmt.Printf("Bandit 1 password: %s", pass1)

	time.Sleep(10)

	pass2 := completeLevel(1, pass1, "whoami")
	fmt.Println("Bandit 1 output: ", pass2)
}

func completeLevel(no int, pass string, cmd string) string {
	nextLevel := fmt.Sprintf("bandit%s", strconv.Itoa(no))
	fmt.Println("Trying level", nextLevel, "using pass", pass)
	nextLevelPass, err := remoteRun(nextLevel, pass, cmd)
	if err != nil {
		fmt.Println("Error executing completeLevel: ", err)
		os.Exit(3)
	}

	return nextLevelPass
}

func getBandit0Password() string {
	requestURL := "https://overthewire.org/wargames/bandit/bandit0.html"
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Println("Error making HTTP request: ", err)
		os.Exit(1)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing http body: ", err)
			os.Exit(1)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		fmt.Println("Didn't receive 200 OK from https://overthewire.org/wargames/bandit/bandit0.html, please investigate.")
		os.Exit(1)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		os.Exit(1)
	}

	strResBody := string(resBody)

	passwordStartSubstring := "password is <strong>"
	passwordStart := strings.Index(strResBody, passwordStartSubstring)
	passwordEnd := strings.Index(strResBody, "</strong>. Once")
	if passwordStart != -1 {
		return strResBody[passwordStart+len(passwordStartSubstring) : passwordEnd]
	}

	return "Not found"
}

func remoteRun(user string, pass string, cmd string) (string, error) {
	// Authentication
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
	}
	// Connect
	client, err := ssh.Dial("tcp", net.JoinHostPort("bandit.labs.overthewire.org", "2220"), config)
	if err != nil {
		return "Error connecting: ", err
	}
	// Create a session. It is one session per command.
	session, err := client.NewSession()
	if err != nil {
		return "Session error: ", err
	}
	defer session.Close()
	var b bytes.Buffer  // import "bytes"
	session.Stdout = &b // get output
	// you can also pass what gets input to the stdin, allowing you to pipe
	// content from client to server
	//      session.Stdin = bytes.NewBufferString("My input")

	// Finally, run the command
	err = session.Run(cmd)
	return b.String(), err
}
