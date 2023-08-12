package main

import (
	"fmt"
	"os"
	"net"
	"github.com/fatih/color"
	"bufio"
    "io/ioutil"	
    "strings"
	getopt "github.com/pborman/getopt/v2"
)
	

func printBanner() {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Printf("\n%s%s\n", yellow("_______       _____      "), red("__________ "))                      
	fmt.Printf("%s%s\n", yellow("___    |___  ___  /_______"), red("__  ____/__________  ________ ___"))                      
	fmt.Printf("%s%s\n", yellow("__  /| |  / / /  __/  __ \\"), red("_  __/  __  __ \\  / / /_  __ `__ \\"))                      
	fmt.Printf("%s%s\n", yellow("_  ___ / /_/ // /_ / /_/ /"), red("  /___  _  / / / /_/ /_  / / / / /"))                      
	fmt.Printf("%s%s\n", yellow("/_/  |_\\__,_/ \\__/ \\____/"), red("/_____/  /_/ /_/\\__,_/ /_/ /_/ /_/ "))                      
	fmt.Printf("%s%s\n\n", green("                    by 0x5ubt13"), cyan("                 Go version"))   
}

func errorMsg(errMsg string) {
	red := color.New(color.FgRed).PrintfFunc()
	red("[-] Error detected: %s\n", errMsg)
}

func checks() {
	// Check 1: If verbose, print banner
	if *optQuiet == false {printBanner()}
	
	// Check 2: Help flag passed?
	if *optHelp {
		if *optQuiet == false {(color.Cyan("[*] Help flag detected. Aborting other checks and printing usage.\n\n"))}
        getopt.Usage()
        os.Exit(0)
    }

	// Check 3: am I groot?!
	if os.Geteuid() != 0 {errorMsg("Please run me as root!")}

	// Check 4: Ensure there is a target
	if *optTarget == "" {
		errorMsg("You must provide an IP address or targets file with the flag -t to start the attack.")
		os.Exit(1)
	}

	// Check 5: Determine whether it is a single target or multi-target   
	targetInput := net.ParseIP(*optTarget)
	if targetInput.To4() == nil {
		// Multi-target, check how many targets are there
		
		// fetching the file
		data, err := ioutil.ReadFile(*optTarget)
		if err != nil {panic(err)}

		// Get lines
		lines := strings.Split(string(data), "\n")
		totalTargetLines := len(lines)
	}

	return totalTargetLines
}
                                                            