package main

import (
	"fmt"

	util "github.com/DevopsGuyXD/IAM-Access-Key-Rotation/Utils"
)

func main() {

	// ====================== Entry Point ========================
	fmt.Printf("\nWelcome to IAM-Notifier\n")
	util.InitAws()
	GetAllIAMUsers()
}