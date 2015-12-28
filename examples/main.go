package main

import "log"

func main() {
	// Check packages
	if err := CheckPackages(); err != nil {
		log.Fatal(err)
	}

	// Check services
	if err := CheckServices(); err != nil {
		log.Fatal(err)
	}

	// Check containers
	if err := CheckContainers(); err != nil {
		log.Fatal(err)
	}

	// Check filesystems
	if err := CheckFiles(); err != nil {
		log.Fatal(err)
	}

	// Check processes
	if err := CheckProcesses(); err != nil {
		log.Fatal(err)
	}
}
