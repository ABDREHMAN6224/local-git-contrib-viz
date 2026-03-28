package main

import (
	"flag"
	"fmt"
)

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()
	if folder != "" {
		scan(folder)
		return
	}
	fmt.Printf("Scanning for email: %s\n\n", email)
	stats(email)
}
