package main

import(
	"fmt"
	"net/http"
	"strings"
	"strconv"
)

func main() {
	websites := []string {
		"https://google.com",
		"https://stackoverflow.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
	}

	// Concurrency using Go Routines and Channels
	c := make(chan string)

	for _, site := range websites {
		go checkSiteStatus(site, c)
	}

	for i := 0; i < len(websites); i++ {
		fmt.Println(<-c)
	}
}

func checkSiteStatus(site string, c chan string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Site ", site, "returned an error: ", err)
		return
	}
	defer resp.Body.Close()
	msg := []string{"Site: ", site, "\t Status Code: ", strconv.Itoa(resp.StatusCode)}
	c <- strings.Join(msg[:], "")
}