package main

import (
	"net/http"
	"fmt"
)

func main() {
	websites := []string {
		"https://google.com",
		"https://stackoverflow.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
	}

	// Secuential using a serial for loop
	for _, site := range websites {
		code := checkSiteStatus(site)
		fmt.Printf("Site: %v\tStatus code: %v\n", site[8:], code)
	}
}

func checkSiteStatus(site string) int {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Site ", site, "returned an error: ", err)
		return 1
	}
	defer resp.Body.Close()

	return resp.StatusCode
}