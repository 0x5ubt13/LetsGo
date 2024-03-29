package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/fatih/color"
	getopt "github.com/pborman/getopt/v2"
)

var (
    yellow 	= color.New(color.FgYellow).SprintFunc()
	red 	= color.New(color.FgRed).SprintFunc()
	green 	= color.New(color.FgGreen).SprintFunc()
	cyan 	= color.New(color.FgCyan).SprintFunc()
)

func printBanner() {
	fmt.Printf("\n%s%s%s\n", yellow("_______       _____      "), red("__________ "), cyan("               Go version"))                     
	fmt.Printf("%s%s\n", yellow("___    |___  ___  /_______"), red("__  ____/__________  ________ ___"))                      
	fmt.Printf("%s%s\n", yellow("__  /| |  / / /  __/  __ \\"), red("_  __/  __  __ \\  / / /_  __ `__ \\"))                      
	fmt.Printf("%s%s\n", yellow("_  ___ / /_/ // /_ / /_/ /"), red("  /___  _  / / / /_/ /_  / / / / /"))                      
	fmt.Printf("%s%s\n", yellow("/_/  |_\\__,_/ \\__/ \\____/"), red("/_____/  /_/ /_/\\__,_/ /_/ /_/ /_/ "))                      
	fmt.Printf("%s\n\n", green("                    by 0x5ubt13"))   

}

// Define a global regular expression pattern
var alphanumericRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

// Use isAlphanumeric for regexp
func isAlphanumeric(s string) bool {
  return alphanumericRegex.MatchString(s)
}


// Custom error message printed out to terminal
func errorMsg(errMsg string) {
	red("[-] Error detected: %s\n", errMsg)
}


// Perform pre-flight checks and return total lines if multi-target
func checks() int {
	var totalLines int

	// Check 1: If verbose, print banner
	if !*optQuiet {printBanner()}
	
	// Check 2: Help flag passed?
	if *optHelp {
		if !*optQuiet {(color.Cyan("[*] Help flag detected. Aborting other checks and printing usage.\n\n"))}
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
	
	// Check 5: Ensure base output directory is correctly set and exists
	customMkdir(*optOutput)

	// Check 6: Determine whether it is a single target or multi-target   
	targetInput := net.ParseIP(*optTarget)
	fmt.Printf("Debug: targetInput := %s\n", targetInput.To4())
	if targetInput.To4() == nil {
		// Multi-target
		// Check file exists and get lines
		_, totalLines = readTargetsFile(*optTarget)
	} else {
		totalLines = 0
	}

	// Check 7: locate exists in the system
	checkProgramExists("locate")

	return totalLines
}


func checkProgramExists(command string) {
	_, err := exec.LookPath(command)
	if err != nil {
		fmt.Println(fmt.Errorf("AutoEnum needs '%s' to be installed. Please install it manually", command))
		os.Exit(1)
	} else {
		fmt.Printf("'%s' is installed.\n", command)
	}
}
		

func IsValidPath(fp string) bool {
	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
	  return true
	}

	// Attempt to create it
	var d []byte
	if err := os.WriteFile(fp, d, 0644); err == nil {
	  os.Remove(fp) // And delete it
	  return true
	}

	return false
}
                  

func readTargetsFile(filename string) ([]string, int) {
	// fetching the file
	data, err := os.ReadFile(*optTarget)
	if err != nil {panic(err)}

	// Get lines
	lines := strings.Split(string(data), "\n")
	return lines, len(lines)-1
}


func printPhase(phase int) {
	if !*optQuiet {
		fmt.Printf("\n%s%s ", cyan("[*] ---------- "), "Starting Phase")
		switch phase {
		case 0:
			fmt.Printf("%s%s", yellow("0"), ": running initial checks ")
		case 1:
			fmt.Printf("%s%s", yellow("1"), ": parsing the CIDR range ")
		case 2:
			fmt.Printf("%s%s", yellow("2"), ": parsing target or list of targets ")
		case 3:
			fmt.Printf("%s%s", yellow("3"), ": parsing found ports ")
		case 33:
			fmt.Printf("%s%s", yellow("3"), ": running multi-target mode. Looping through the list, one target at a time ")
		case 4:
			fmt.Printf("%s%s", yellow("4"), ": background tools working ")
		default:
			errorMsg("Development error. There are currently 5 phases in the script ")
		}
		fmt.Printf("%s\n\n", cyan("----------"))
	}
}


func customMkdir(name string) {
	if IsValidPath(name){
		err := os.Mkdir(name, os.ModePerm)
		if err != nil {
			fmt.Println(red("[-] Error:"), red(err))
		} else {
			fmt.Printf("%s %s %s", green("[+] Directory"), yellow(name), green("created successfully"))
		}
	}
}


// Write bytes output to file
func writePortsToFile(filePath string, ports string, host string) string {
	// Open file
	fileName := fmt.Sprintf("%s/open_ports.txt", filePath)
	f, err := os.Create(fileName)

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

	_, err2 := fmt.Fprintln(f, ports)

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Printf("%s %s %s %s\n", green("[+] Successfully written open ports for host"), yellow(host), green("to file"), yellow(fileName))

	return ports
}
