package main

import (
    "context"
    "fmt"
    "log"
    "time"
	"os"

	getopt "github.com/pborman/getopt/v2"
    nmap "github.com/Ullaakut/nmap/v3"
)

func main() {
	// getopts
    optAgain := getopt.BoolLong("again", 'a', 0, "Repeat the scan and compare with initial ports discovered.")
    optName := getopt.StringLong("name", 'n', "", "Your name")
    optName := getopt.StringLong("name", 'n', "", "Your name")
    optName := getopt.StringLong("name", 'n', "", "Your name")
    optName := getopt.StringLong("name", 'n', "", "Your name")
    optName := getopt.StringLong("name", 'n', "", "Your name")
    optName := getopt.StringLong("name", 'n', "", "Your name")
    optName := getopt.StringLong("name", 'n', "", "Your name")
    optName := getopt.StringLong("name", 'n', "", "Your name")
    optHelp := getopt.BoolLong("help", 0, "Help")
    getopt.Parse()

    if *optHelp {
        getopt.Usage()
        os.Exit(0)
    }

	fmt.Println("Hello " + *optName + "!")

	// scan
	scan()
}

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