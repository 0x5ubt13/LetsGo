package main

import (
	"context"
	"fmt"
	"log"
	// "os"
	"time"

	// "github.com/fatih/color"
	nmap "github.com/Ullaakut/nmap/v3"
	// "github.com/fatih/color"
	getopt "github.com/pborman/getopt/v2"
	
)

// Declare flags and have getopt return pointers to the values.
var optAgain 	= getopt.BoolLong("again", 'a', "Repeat the scan and compare with initial ports discovered.")
var optBrute	= getopt.BoolLong("brute", 'b', "Activate all fuzzing and bruteforcing in the script.")
var optDNS 		= getopt.StringLong("DNS", 'd', "", "Specify custom DNS servers. Default option: -n")
var optHelp 	= getopt.BoolLong("help", 'h', "Display this help and exit.")
var optTopPorts = getopt.StringLong("top", 'p', "", "Run port sweep with nmap and the flag --top-ports=<your input>")
var optQuiet 	= getopt.BoolLong("quiet", 'q', "Don't print the banner and decrease overall verbosity.")
var optRange 	= getopt.StringLong("range", 'r', "", "Specify a CIDR range to use tools for whole subnets")
var optSlower 	= getopt.BoolLong("slower", 's', "Don't use Rustscan for the initial port sweep")
var optTarget 	= getopt.StringLong("target", 't', "", "Specify target single IP / List of IPs file.")

func main() {
	// Parse the program arguments
	getopt.Parse()
	// Get the remaining positional parameters
	// args := getopt.Args()

	fmt.Println("Debug:")
	fmt.Printf("Again: %t\n", *optAgain)
	fmt.Printf("Brute: %t\n", *optBrute)
	fmt.Printf("DNS: %s\n", *optDNS)
    fmt.Printf("Help: %t\n", *optHelp) 	
    fmt.Printf("Top ports: %s\n", *optTopPorts) 
    fmt.Printf("Quiet: %t\n", *optQuiet)	
    fmt.Printf("Range: %s\n", *optRange)	
    fmt.Printf("Slower: %t\n", *optSlower)	
    fmt.Printf("Target: %s\n", *optTarget)	
	
	// Checks
	checks()

	// scan
	//scan()

	fmt.Println("End of main function")
}


// Run scan
func scan() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5-minute timeout.
	scanner, err := nmap.NewScanner(
		ctx,
		nmap.WithTargets("google.com", "facebook.com", "youtube.com"),
		nmap.WithPorts("80,443,843"),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		log.Printf("run finished with warnings: %s\n", *warnings) // Warnings are non-critical errors from nmap.
	}
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	fmt.Printf("Nmap done: %d hosts up scanned in %.2f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
}