package main

import (
	"os"
	"os/exec"
	"fmt"
)

func main() {
	
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	fmt.Printf("arg 0: %s\n", argsWithProg)
    fmt.Printf("arg 1: %s\n", argsWithoutProg)

	// ip := os.Args[1]

	nmap_allports := exec.Command("nmap", "-p- --min-rate=1000 -T4 %s", os.Args[1], "-Pn" )
	
	stdout, err := nmap_allports.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))

}