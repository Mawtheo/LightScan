package main

// TODO : Faire un clear du terminal avant de lancer l'app en fonction de l'OS
// TODO : Vérifier Port soit correct lors de l'input de l'user
// TODO : Utiliser goroutines
// TODO : Séparer les fichiers
// TODO : Faire une interface web (VueJS, shadcn, Gin)
// TODO : Scan de ports connus (port 80, 443, 22, etc.)
// TODO : Faire des TESTS

import (
	"LightScan/scan"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func scanPort() {
	var ip string
	var firstPort int
	var lastPort int

	fmt.Println("")
	fmt.Println("Enter the IP address to scan")
	fmt.Print("IP: ")
	fmt.Scan(&ip)
	// Verify IP is valid
	if net.ParseIP(ip) == nil {
		fmt.Println("Invalid IP address format !")
		os.Exit(1)
	}
	fmt.Println("")

	fmt.Println("Select the first port to scan")
	fmt.Print("First port: ")
	fmt.Scan(&firstPort)
	// Check port
	if !verifyPort(firstPort) {
		os.Exit(1)
	}
	fmt.Println("")

	fmt.Println("Select the last port to scan")
	fmt.Print("Last port: ")
	fmt.Scan(&lastPort)
	// Check port
	if !verifyPort(lastPort) {
		os.Exit(1)
	}
	// Check Ports range
	switch {
	case firstPort == lastPort:
		fmt.Println("The first port is the same as the last port !")
		os.Exit(1)
	case firstPort > lastPort:
		fmt.Println("Cannot be greater than the first port !")
		os.Exit(1)
	}
	fmt.Println("")

	for port := firstPort; port <= lastPort; port++ {
		if scanEachPort(ip, port) {
			portName, exist := scan.KnownPorts[port]
			if exist {
				fmt.Printf("Port: %d (%s) is UP !\n", port, portName)
			} else {
				fmt.Println("Port:", port, "is UP !")
			}
		} else {
			fmt.Println("Port:", port)
		}
	}
}

func scanEachPort(ip string, port int) bool {
	// Format example: {192.168.1.100}:{80}
	address := ip + ":" + strconv.Itoa(port)
	// Set the time for each connection
	var timeout time.Duration = 1 * time.Second
	// TCP check if port is open during timeout
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		// Port is close
		return false
	}
	// Close connection
	conn.Close()
	// Port is open
	return true
}

func verifyPort(port int) bool {
	// Verify port in range uint16
	if port < 0 || port > 65535 {
		fmt.Println("Must be in range [0:65535] !")
		return false
	}
	return true
}

func main() {
	fmt.Println("")
	fmt.Println("\033[31m // LightScan // \033[0m")
	fmt.Println("\033[31m // Scan any TCP port(s) is open !// \033[0m")
	fmt.Println("")

	scanPort()
}
