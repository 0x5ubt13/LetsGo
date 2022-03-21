package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	content := readSite(url)
	// fmt.Println(content)	

	tours := parseJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name + ": $" + tour.Price)
	}
}

func parseJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)

		tours = append(tours, tour)
	}

	return tours
}

type Tour struct {
	Name, Price string
}

func readSite(url string) string {
	resp, err := http.Get(url)
	checkError(err)
	//fmt.Printf("Response type: %T\n", resp) // output = "Response type: *http.Response"

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}