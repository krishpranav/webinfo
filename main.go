package main

import (
	"fmt"
	"sync"

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
	fmt.Println("usage: webinfo subcommand { options }")
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

func examples() {
	fmt.Println("	Examples:")
	fmt.Println("		- webinfo dns -target target.domain")
	fmt.Println("		- webinfo dns -target -o txt target.domain")
	fmt.Println("		- webinfo dns -target -o html target.domain")
	fmt.Println("		- webinfo dns -target -plain target.domain")
	fmt.Println()
	fmt.Println("		- webinfo subdomain -target target.domain")
	fmt.Println("		- webinfo subdomain -w wordlist.txt -target target.domain")
	fmt.Println("		- webinfo subdomain -o txt -target target.domain")
	fmt.Println("		- webinfo subdomain -o html -target target.domain")
	fmt.Println("		- webinfo subdomain -i 400 -target target.domain")
	fmt.Println("		- webinfo subdomain -i 4** -target target.domain")
	fmt.Println("		- webinfo subdomain -c -target target.domain")
	fmt.Println("		- webinfo subdomain -db -target target.domain")
	fmt.Println("		- webinfo subdomain -plain -target target.domain")
	fmt.Println()
	fmt.Println("		- webinfo port -p -450 -target target.domain")
	fmt.Println("		- webinfo port -p 90- -target target.domain")
	fmt.Println("		- webinfo port -p 10-1000 -target target.domain")
	fmt.Println("		- webinfo port -o txt -target target.domain")
	fmt.Println("		- webinfo port -o html -target target.domain")
	fmt.Println("		- webinfo port -p 21,25,80 -target target.domain")
	fmt.Println("		- webinfo port -common -target target.domain")
	fmt.Println("		- webinfo port -plain -target target.domain")
	fmt.Println()
	fmt.Println("		- webinfo dir -target target.domain")
	fmt.Println("		- webinfo dir -o txt -target target.domain")
	fmt.Println("		- webinfo dir -o html -target target.domain")
	fmt.Println("		- webinfo dir -w wordlist.txt -target target.domain")
	fmt.Println("		- webinfo dir -i 500,401 -target target.domain")
	fmt.Println("		- webinfo dir -i 5**,401 -target target.domain")
	fmt.Println("		- webinfo dir -c -target target.domain")
	fmt.Println("		- webinfo dir -plain -target target.domain")
	fmt.Println()
	fmt.Println("		- webinfo report -p 80 -target target.domain")
	fmt.Println("		- webinfo report -o txt -target target.domain")
	fmt.Println("		- webinfo report -o html -target target.domain")
	fmt.Println("		- webinfo report -p 50-200 -target target.domain")
	fmt.Println("		- webinfo report -wd dirs.txt -target target.domain")
	fmt.Println("		- webinfo report -ws subdomains.txt -target target.domain")
	fmt.Println("		- webinfo report -id 500,501,502 -target target.domain")
	fmt.Println("		- webinfo report -is 500,501,502 -target target.domain")
	fmt.Println("		- webinfo report -id 5**,4** -target target.domain")
	fmt.Println("		- webinfo report -is 5**,4** -target target.domain")
	fmt.Println("		- webinfo report -cd -target target.domain")
	fmt.Println("		- webinfo report -cs -target target.domain")
	fmt.Println("		- webinfo report -db -target target.domain")
	fmt.Println("		- webinfo report -p 21,25,80 -target target.domain")
	fmt.Println("		- webinfo report -common -target target.domain")
	fmt.Println("")
}

func main() {
	input := readArgs()
	subs := make(map[string]Asset)
	dirs := make(map[string]Asset)

	common := []int{13, 20, 21, 22, 23, 25, 42, 50, 51, 53, 67, 68,
		69, 70, 79, 80, 88, 102, 107, 109, 110, 111, 113, 115, 118,
		119, 123, 135, 136, 137, 138, 139, 143, 156, 161, 162, 179,
		194, 220, 311, 389, 443, 445, 464, 500, 512, 513, 514, 515,
		530, 543, 546, 547, 556, 587, 631, 636, 660, 749, 802, 853,
		873, 902, 989, 990, 992, 993, 994, 995, 1000, 1025, 1080,
		1194, 1241, 1293, 1337, 1417, 1433, 1434, 1527, 1755, 1812,
		1813, 1880, 1883, 2000, 2049, 2095, 2096, 2222, 2483, 2484,
		2638, 3000, 3268, 3283, 3333, 3306, 3389, 4000, 4444, 5000,
		5432, 5555, 5938, 6000, 6666, 7000, 7071, 7777, 8000, 8001,
		8002, 8003, 8004, 8005, 8080, 8200, 8888, 9050, 10000}
	execute(input, subs, dirs, common)
}

type Asset struct {
	Value   string
	Printed bool
}

func execute(input Input, subs map[string]Asset, dirs map[string]Asset, common []int) {

	var mutex = &sync.Mutex{}
	if input.ReportTarget != "" {
		intro()
		target := cleanProtocol(input.ReportTarget)
		var targetIP string
		fmt.Printf("target: %s\n", target)

	}
}
