package main

import (
	"fmt"

	"github.com/fatih/color"
)

func intro() {
	banner := "WEB INFO GATHER"
	banner2 := "> github.com/krishpranav/webinfo"
	banner3 := "> Author: krishpranav"
	bannerPart1 := banner
	bannerPart2 := banner2 + banner3
	color.Cyan("%s\n", bannerPart1)
	fmt.Println(bannerPart2)
	fmt.Println("================================")
}

// help function for printing the usage of web info
func help() {
	fmt.Println("Information Gathering tool - DNS / Subdomain / Ports / Directories enumeration")
	fmt.Println("")
	fmt.Println("usage: scilla subcommand { options }")
	fmt.Println("")
	fmt.Println("   Available subcommands:")
	fmt.Println("       - dns [-o output-format]")
	fmt.Println("             [-plain Print only results]")
	fmt.Println("             -target <target (URL/IP)> REQUIRED")
	fmt.Println("       - port [-p <start-end> or ports divided by comma]")
	fmt.Println("              [-o output-format]")
	fmt.Println("              [-common scan common ports]")
	fmt.Println("              [-plain Print only results]")
	fmt.Println("              -target <target (URL/IP)> REQUIRED")
	fmt.Println("       - subdomain [-w wordlist]")
	fmt.Println("                   [-o output-format]")
	fmt.Println("                   [-i ignore status codes]")
	fmt.Println("                   [-c use also a web crawler]")
	fmt.Println("                   [-db use also a public database]")
	fmt.Println("                   [-plain Print only results]")
	fmt.Println("                   -target <target (URL)> REQUIRED")
	fmt.Println("       - dir [-w wordlist]")
	fmt.Println("             [-o output-format]")
	fmt.Println("             [-i ignore status codes]")
	fmt.Println("             [-c use also a web crawler]")
	fmt.Println("             [-plain Print only results]")
	fmt.Println("             -target <target (URL)> REQUIRED")
	fmt.Println("       - report [-p <start-end> or ports divided by comma]")
	fmt.Println("                [-ws subdomains wordlist]")
	fmt.Println("                [-wd directories wordlist]")
	fmt.Println("                [-o output-format]")
	fmt.Println("                [-id ignore status codes in directories scanning]")
	fmt.Println("                [-is ignore status codes in subdomains scanning]")
	fmt.Println("                [-cd use also a web crawler for directories scanning]")
	fmt.Println("                [-cs use also a web crawler for subdomains scanning]")
	fmt.Println("                [-db use also a public database for subdomains scanning]")
	fmt.Println("                [-common scan common ports]")
	fmt.Println("                -target <target (URL/IP)> REQUIRED")
	fmt.Println("       - help")
	fmt.Println("       - examples")
	fmt.Println()
}
