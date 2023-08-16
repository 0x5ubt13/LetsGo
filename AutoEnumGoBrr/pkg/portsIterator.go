package main

import (
	"fmt"
	"strings"
)

// Core functionality of the script
// Iterate through each port and automate launching tools
func portsIterator(targetIP string, baseDir string, openPorts string) {
	// ports covered so far: "21,22,25,465,587,53,79,80,443,8080,88,110,143,993,995,111,137,138,139,445,161,162,623,873,1433,1521"
	portsArray := strings.Split(openPorts, ",")

	// Iterate over each port in the array
	for _, port := range portsArray {
		switch port {
		case "21":
			// Handle port 21
			fmt.Printf("[+] FTP service detected. Running FTP nmap enum scripts.")
			ftpDir := baseDir + "ftp/"
			customMkdir(ftpDir)

			// Running Nmap scripts for FTP
			cmd := exec.Command("nmap", "-sV", "-n", "-Pn", "-p21", targetIP, "--script", "ftp-* and not brute", "-v")
			nmapOutput, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error running nmap:", err)
			}
			nmapOutputFile := ftpDir + "ftp_enum.nmap"
			err = os.WriteFile(nmapOutputFile, nmapOutput, 0644)
			if err != nil {
				fmt.Println("Error writing nmap output:", err)
			}

			// Running Hydra bruteforcing for FTP
			// Implement hydraBruteforcing function or similar here
			// hydraBruteforcing(targetIP, ftpDir, "ftp")
        case "22":
            fmt.Println(" - Secure Shell")
			fmt.Println("[+] SSH service detected. Running SSH nmap enum scripts.")
        case "25":
            fmt.Println("SMTP - Simple Mail Transfer Protocol")
        case "465":
            fmt.Println("SMTPS - Simple Mail Transfer Protocol Secure")
        case "587":
            fmt.Println("SMTP - Message Submission")
        case "53":
            fmt.Println("DNS - Domain Name System")
        case "79":
            fmt.Println("Finger")
        case "80":
            fmt.Println("HTTP - HyperText Transfer Protocol")
        case "443":
            fmt.Println("HTTPS - HyperText Transfer Protocol Secure")
        case "8080":
            fmt.Println("HTTP - Alternate Port")
        case "88":
            fmt.Println("Kerberos")
        case "110":
            fmt.Println("POP3 - Post Office Protocol")
        case "143":
            fmt.Println("IMAP - Internet Message Access Protocol")
        case "993":
            fmt.Println("IMAP - Secure")
        case "995":
            fmt.Println("POP3 - Secure")
        case "111":
            fmt.Println("RPC - Remote Procedure Control")
        case "137":
            fmt.Println("NetBIOS Name Service")
        case "138":
            fmt.Println("NetBIOS Datagram Service")
        case "139":
            fmt.Println("NetBIOS Session Service")
        case "445":
            fmt.Println("SMB - Server Message Block")
        case "161":
            fmt.Println("SNMP - Simple Network Management Protocol")
        case "162":
            fmt.Println("SNMP - Trap")
        case "623":
            fmt.Println("ASF - Remote Management and Control Protocol")
        case "873":
            fmt.Println("RSYNC - Remote Sync")
        case "1433":
            fmt.Println("Microsoft SQL Server")
        case "1521":
            fmt.Println("Oracle Database")
        default:
            fmt.Println("Unknown port")
        }
	}

}