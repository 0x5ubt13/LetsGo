package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	nmap "github.com/Ullaakut/nmap/v3"
)

func portSweep(target string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5-minute timeout.
	scanner, err := nmap.NewScanner(
		ctx,
		nmap.WithTargets(target),
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


func singleTarget(target string, baseFilePath string) {
	fmt.Printf("Debug - Single target: %s\n", target)
	fmt.Printf("Debug - Base file path: %s\n", baseFilePath)
	
	path := fmt.Sprintf("%s/%s", baseFilePath, target)
	// Make base dir for the target
	customMkdir(path)

	portSweep(target)
}


func multiTarget(targetsFile *string) {
	if !*optQuiet { fmt.Printf("%s%s\n", green("[+] Using multi-targets file: "), yellow(*targetsFile)) }
	fileNameWithExtension := strings.Split(*targetsFile, "/")
	fileName := strings.Split(fileNameWithExtension[len(fileNameWithExtension)-1], ".")

	// Make base folder for the output
	targetsBaseFilePath := fmt.Sprintf("%s/%s", *optOutput, fileName[0])
	customMkdir(targetsBaseFilePath)

	// Loop through the targets in the file
	targets, lines := readTargetsFile(*targetsFile)
	if !*optQuiet { fmt.Printf("%s %s %s\n", green("[+] Found"), yellow(lines), green("targets")) }
	for i := 0; i < lines; i++ {
		target := targets[i]
		fmt.Printf("%s %v %s %v: %s\n", green("[+] Attacking target"), yellow(i+1), green("of"), yellow(lines), yellow(target))
		singleTarget(target, targetsBaseFilePath)
	}
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