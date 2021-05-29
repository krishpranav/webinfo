package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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
		fmt.Println("=============== FULL REPORT ===============")
		outputFile := ""
		if input.ReportOutput != "" {
			outputFile = createOutputFile(input.ReportTarget, "report", input.ReportOutput)
			if outputFile[len(outputFile)-4:] == "html" {
				bannerHTML(input.ReportTarget, outputFile)
			}
		}

		fmt.Println("=============== SUBDOMAINS SCANNING ===============")
		var strings1 []string

		if isIP(target) {
			targetIP = target
			target = ipToHostname(target)
		}
		if outputFile != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				headerHTML("SUBDOMAIN SCANNING", outputFile)
			}
		}
		if input.ReportCrawlerSub {
			go spawnCrawler(target, input.ReportIgnoreSub, dirs, subs, outputFile, mutex, "sub", false)
		}
		strings1 = createSubdomains(input.ReportWordSub, target)
		if input.ReportSubdomainDB {
			sonar := sonarSubdomains(target, false)
			strings1 = appendDBSubdomains(sonar, strings1)
			crtsh := crtshSubdomains(target, false)
			strings1 = appendDBSubdomains(crtsh, strings1)
			threatcrowd := threatcrowdSubdomains(target, false)
			strings1 = appendDBSubdomains(threatcrowd, strings1)
			hackerTarget := hackerTargetSubdomains(target, false)
			strings1 = appendDBSubdomains(hackerTarget, strings1)
			bufferOverrun := bufferOverrunSubdomains(target, false)
			strings1 = appendDBSubdomains(bufferOverrun, strings1)
			fmt.Fprint(os.Stdout, "\r \r \r \r")
		}

		strings1 = removeDuplicateValues(strings1)
		asyncGet(strings1, input.ReportIgnoreSub, outputFile, subs, mutex, false)
		if outputFile != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				footerHTML(outputFile)
			}
		}

		if targetIP != "" {
			target = targetIP
		}
		fmt.Println("=============== PORT SCANNING ===============")

		asyncPort(input.portsArray, input.portArrayBool, input.StartPort, input.EndPort, target, outputFile, input.ReportCommon, common, false)

		fmt.Println("=============== DNS SCANNING ===============")
		lookupDNS(target, outputFile, false)

		fmt.Println("=============== DIRECTORIES SCANNING ===============")
		var strings2 []string
		strings2 = createUrls(input.ReportWordDir, target)
		if outputFile != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				headerHTML("DIRECTORY SCANNING", outputFile)
			}
		}
		if input.ReportCrawlerDir {
			go spawnCrawler(target, input.ReportIgnoreDir, dirs, subs, outputFile, mutex, "dir", false)
		}
		asyncDir(strings2, input.ReportIgnoreDir, outputFile, dirs, mutex, false)
		if outputFile != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				footerHTML(outputFile)
			}
		}
		if input.ReportOutput != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				bannerFooterHTML(outputFile)
			}
		}
	}

	if input.DNSTarget != "" {
		if !input.DNSPlain {
			intro()
		}
		target := cleanProtocol(input.DNSTarget)
		if isIP(target) {
			target = ipToHostname(target)
		}
		if !input.DNSPlain {
			fmt.Printf("target: %s\n", target)
			fmt.Println("=============== DNS SCANNING ===============")
		}
		outputFile := ""
		if input.DNSOutput != "" {
			outputFile = createOutputFile(input.DNSTarget, "dns", input.DNSOutput)

			if outputFile[len(outputFile)-4:] == "html" {
				bannerHTML(input.DNSTarget, outputFile)
			}
		}
		lookupDNS(target, outputFile, input.DNSPlain)
		if input.DNSOutput != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				bannerFooterHTML(outputFile)
			}
		}
	}

	if input.SubdomainTarget != "" {

		if !input.SubdomainPlain {
			intro()
		}

		target := cleanProtocol(input.SubdomainTarget)
		if isIP(target) {
			target = ipToHostname(target)
		}
		if !input.SubdomainPlain {
			fmt.Printf("target: %s\n", target)
			fmt.Println("=============== SUBDOMAINS SCANNING ===============")
		}
		outputFile := ""
		if input.SubdomainOutput != "" {
			outputFile = createOutputFile(input.SubdomainTarget, "subdomain", input.SubdomainOutput)
			if outputFile[len(outputFile)-4:] == "html" {
				bannerHTML(input.SubdomainTarget, outputFile)
			}
		}
		var strings1 []string
		strings1 = createSubdomains(input.SubdomainWord, target)
		if input.SubdomainDB {
			sonar := sonarSubdomains(target, input.SubdomainPlain)
			strings1 = appendDBSubdomains(sonar, strings1)
			crtsh := crtshSubdomains(target, input.SubdomainPlain)
			strings1 = appendDBSubdomains(crtsh, strings1)
			threatcrowd := threatcrowdSubdomains(target, input.SubdomainPlain)
			strings1 = appendDBSubdomains(threatcrowd, strings1)
			hackerTarget := hackerTargetSubdomains(target, input.SubdomainPlain)
			strings1 = appendDBSubdomains(hackerTarget, strings1)
			bufferOverrun := bufferOverrunSubdomains(target, input.SubdomainPlain)
			strings1 = appendDBSubdomains(bufferOverrun, strings1)
			if !input.SubdomainPlain {
				fmt.Fprint(os.Stdout, "\r \r \r \r")
			}
		}
		if outputFile != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				headerHTML("SUBDOMAIN SCANNING", outputFile)
			}
		}
		if input.SubdomainCrawler {
			go spawnCrawler(target, input.SubdomainIgnore, dirs, subs, outputFile, mutex, "sub", input.SubdomainPlain)
		}

		strings1 = removeDuplicateValues(strings1)
		asyncGet(strings1, input.SubdomainIgnore, outputFile, subs, mutex, input.SubdomainPlain)
		if outputFile != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				footerHTML(outputFile)
			}
		}
		if input.SubdomainOutput != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				bannerFooterHTML(outputFile)
			}
		}
	}

	if input.DirTarget != "" {

		if !input.DirPlain {
			intro()
		}

		target := cleanProtocol(input.DirTarget)
		if !input.DirPlain {
			fmt.Printf("target: %s\n", target)
			fmt.Println("=============== DIRECTORIES SCANNING ===============")
		}
		outputFile := ""
		if input.DirOutput != "" {
			outputFile = createOutputFile(input.DirTarget, "dir", input.DirOutput)

			if outputFile[len(outputFile)-4:] == "html" {
				bannerHTML(input.DirTarget, outputFile)
			}
		}
		var strings2 []string
		strings2 = createUrls(input.DirWord, target)
		if outputFile != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				headerHTML("DIRECTORY SCANNING", outputFile)
			}
		}
		if input.DirCrawler {
			go spawnCrawler(target, input.DirIgnore, dirs, subs, outputFile, mutex, "dir", input.DirPlain)
		}
		asyncDir(strings2, input.DirIgnore, outputFile, dirs, mutex, input.DirPlain)
		if outputFile != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				footerHTML(outputFile)
			}
		}
		if input.DirOutput != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				bannerFooterHTML(outputFile)
			}
		}
	}

	if input.PortTarget != "" {
		if !input.PortPlain {
			intro()
		}
		target := input.PortTarget
		if isURL(target) {
			target = cleanProtocol(input.PortTarget)
		}
		outputFile := ""
		if input.PortOutput != "" {
			outputFile = createOutputFile(input.PortTarget, "port", input.PortOutput)
			if outputFile[len(outputFile)-4:] == "html" {
				bannerHTML(input.PortTarget, outputFile)
			}
		}
		if !input.PortPlain {
			fmt.Printf("target: %s\n", target)
			fmt.Println("=============== PORT SCANNING ===============")
		}
		asyncPort(input.portsArray, input.portArrayBool, input.StartPort, input.EndPort, target, outputFile, input.PortCommon, common, input.PortPlain)

		if input.PortOutput != "" {
			if outputFile[len(outputFile)-4:] == "html" {
				bannerFooterHTML(outputFile)
			}
		}
	}
}

// clean protocol connections
func cleanProtocol(target string) string {
	if len(target) > 6 {
		// clean protocols and go ahead
		if target[:6] == "tls://" {
			target = target[6:]
		}
	}
	if len(target) > 7 {
		if target[:7] == "http://" {
			target = target[7:]
		}
	}
	if len(target) > 8 {
		if target[:8] == "https://" {
			target = target[8:]
		}
	}
	if target[len(target)-1:] == "/" {
		return target[:len(target)-1]
	}
	return target
}

func outputFormatIsOk(input string) bool {
	if input == "" {
		return true
	}
	acceptedOutput := [2]string{"txt", "html"}
	input = strings.ToLower(input)
	for _, output := range acceptedOutput {
		if output == input {
			return true
		}
	}
	return false
}

type Input struct {
	ReportTarget      string
	ReportWordDir     string
	ReportWordSub     string
	ReportOutput      string
	ReportIgnoreDir   []string
	ReportIgnoreSub   []string
	ReportCrawlerDir  bool
	ReportCrawlerSub  bool
	ReportSubdomainDB bool
	ReportCommon      bool
	DNSTarget         string
	DNSOutput         string
	DNSPlain          bool
	SubdomainTarget   string
	SubdomainWord     string
	SubdomainOutput   string
	SubdomainIgnore   []string
	SubdomainCrawler  bool
	SubdomainDB       bool
	SubdomainPlain    bool
	DirTarget         string
	DirWord           string
	DirOutput         string
	DirIgnore         []string
	DirCrawler        bool
	DirPlain          bool
	PortTarget        string
	PortOutput        string
	StartPort         int
	EndPort           int
	portArrayBool     bool
	portsArray        []int
	PortCommon        bool
	PortPlain         bool
}

func readArgs() Input {
	reportCommand := flag.NewFlagSet("report", flag.ExitOnError)
	dnsCommand := flag.NewFlagSet("dns", flag.ExitOnError)
	subdomainCommand := flag.NewFlagSet("subdomain", flag.ExitOnError)
	portCommand := flag.NewFlagSet("port", flag.ExitOnError)
	dirCommand := flag.NewFlagSet("dir", flag.ExitOnError)
	helpCommand := flag.NewFlagSet("help", flag.ExitOnError)
	examplesCommand := flag.NewFlagSet("examples", flag.ExitOnError)

	reportTargetPtr := reportCommand.String("target", "", "Target {URL/IP} (Required)")

	reportPortsPtr := reportCommand.String("p", "", "ports range <start-end>")

	reportWordlistDirPtr := reportCommand.String("wd", "", "wordlist to use for directories (default enabled)")

	reportWordlistSubdomainPtr := reportCommand.String("ws", "", "wordlist to use for subdomains (default enabled)")

	reportOutputPtr := reportCommand.String("o", "", "output format (txt/html)")

	reportIgnoreDirPtr := reportCommand.String("id", "", "Ignore response code(s) for directories scanning")
	reportIgnoreDir := []string{}

	reportIgnoreSubPtr := reportCommand.String("is", "", "Ignore response code(s) for subdomains scanning")
	reportIgnoreSub := []string{}

	reportCrawlerDirPtr := reportCommand.Bool("cd", false, "Use also a web crawler for directories enumeration")

	reportCrawlerSubdomainPtr := reportCommand.Bool("cs", false, "Use also a web crawler for subdomains enumeration")

	reportSubdomainDBPtr := reportCommand.Bool("db", false, "Use also a public database for subdomains enumeration")

	reportCommonPtr := reportCommand.Bool("common", false, "Scan common ports")

	dnsTargetPtr := dnsCommand.String("target", "", "Target {URL/IP} (Required)")

	dnsOutputPtr := dnsCommand.String("o", "", "output format (txt/html)")

	dnsPlainPtr := dnsCommand.Bool("plain", false, "Print only results")

	subdomainTargetPtr := subdomainCommand.String("target", "", "Target {URL} (Required)")

	subdomainWordlistPtr := subdomainCommand.String("w", "", "wordlist to use (default enabled)")

	subdomainOutputPtr := subdomainCommand.String("o", "", "output format (txt/html)")

	subdomainIgnorePtr := subdomainCommand.String("i", "", "Ignore response code(s)")
	subdomainIgnore := []string{}

	subdomainCrawlerPtr := subdomainCommand.Bool("c", false, "Use also a web crawler")

	subdomainDBPtr := subdomainCommand.Bool("db", false, "Use also a public database")

	subdomainPlainPtr := subdomainCommand.Bool("plain", false, "Print only results")

	dirTargetPtr := dirCommand.String("target", "", "Target {URL/IP} (Required)")

	dirWordlistPtr := dirCommand.String("w", "", "wordlist to use (default enabled)")

	dirOutputPtr := dirCommand.String("o", "", "output format (txt/html)")

	dirIgnorePtr := dirCommand.String("i", "", "Ignore response code(s)")
	dirIgnore := []string{}

	dirCrawlerPtr := dirCommand.Bool("c", false, "Use also a web crawler")

	dirPlainPtr := dirCommand.Bool("plain", false, "Print only results")

	portTargetPtr := portCommand.String("target", "", "Target {URL/IP} (Required)")

	portOutputPtr := portCommand.String("o", "", "output format (txt/html)")

	portsPtr := portCommand.String("p", "", "ports range <start-end>")

	portCommonPtr := portCommand.Bool("common", false, "Scan common ports")

	portPlainPtr := portCommand.Bool("plain", false, "Print only results")

	StartPort := 1
	EndPort := 65535
	portsArray := []int{}
	portArrayBool := false

	if len(os.Args) < 2 {
		intro()
		fmt.Println("[ERROR] subcommand is required.")
		fmt.Println("	Type: scilla help      - Full overview of the commands.")
		fmt.Println("	Type: scilla examples  - Some explanatory examples.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "report":
		reportCommand.Parse(os.Args[2:])
	case "dns":
		dnsCommand.Parse(os.Args[2:])
	case "subdomain":
		subdomainCommand.Parse(os.Args[2:])
	case "port":
		portCommand.Parse(os.Args[2:])
	case "dir":
		dirCommand.Parse(os.Args[2:])
	case "help":
		intro()
		helpCommand.Parse(os.Args[2:])
	case "examples":
		intro()
		examplesCommand.Parse(os.Args[2:])
	default:
		intro()
		flag.PrintDefaults()
		os.Exit(1)
	}

	if reportCommand.Parsed() {

		if *reportTargetPtr == "" {
			reportCommand.PrintDefaults()
			os.Exit(1)
		}

		if !isURL(*reportTargetPtr) {
			fmt.Println("The inputted target is not valid.")
			os.Exit(1)
		}
		if !outputFormatIsOk(*reportOutputPtr) {
			fmt.Println("The output format is not valid.")
			os.Exit(1)
		}

		if *reportPortsPtr != "" && *reportCommonPtr {
			fmt.Println("You can't specify a port range and common option together.")
			os.Exit(1)
		}

		if *reportPortsPtr != "" {
			if strings.Contains(*reportPortsPtr, "-") && strings.Contains(*reportPortsPtr, ",") {
				fmt.Println("You can specify a ports range or an array, not both.")
				os.Exit(1)
			}
			if strings.Contains(*reportPortsPtr, "-") {
				portsRange := string(*reportPortsPtr)
				StartPort, EndPort = checkPortsRange(portsRange, StartPort, EndPort)
				portArrayBool = false
			} else if strings.Contains(*reportPortsPtr, ",") {
				portsArray = checkPortsArray(*reportPortsPtr)
				portArrayBool = true
			} else {
				portsRange := string(*reportPortsPtr)
				StartPort, EndPort = checkPortsRange(portsRange, StartPort, EndPort)
				portArrayBool = false
			}
		}
		if *reportIgnoreDirPtr != "" {
			toBeIgnored := string(*reportIgnoreDirPtr)
			reportIgnoreDir = checkIgnore(toBeIgnored)
		}
		if *reportIgnoreSubPtr != "" {
			toBeIgnored := string(*reportIgnoreSubPtr)
			reportIgnoreSub = checkIgnore(toBeIgnored)
		}
	}

	if dnsCommand.Parsed() {

		if *dnsTargetPtr == "" {
			dnsCommand.PrintDefaults()
			os.Exit(1)
		}
		if !isURL(*dnsTargetPtr) {
			fmt.Println("The inputted target is not valid.")
			os.Exit(1)
		}
		if !outputFormatIsOk(*dnsOutputPtr) {
			fmt.Println("The output format is not valid.")
			os.Exit(1)
		}
	}

	if subdomainCommand.Parsed() {

		if *subdomainTargetPtr == "" {
			subdomainCommand.PrintDefaults()
			os.Exit(1)
		}
		if !isURL(*subdomainTargetPtr) {
			fmt.Println("The inputted target is not valid.")
			os.Exit(1)
		}
		if !outputFormatIsOk(*subdomainOutputPtr) {
			fmt.Println("The output format is not valid.")
			os.Exit(1)
		}
		if *subdomainIgnorePtr != "" {
			toBeIgnored := string(*subdomainIgnorePtr)
			subdomainIgnore = checkIgnore(toBeIgnored)
		}
	}

	if portCommand.Parsed() {
		if *portTargetPtr == "" {
			portCommand.PrintDefaults()
			os.Exit(1)
		}

		if *portsPtr != "" && *portCommonPtr {
			fmt.Println("You can't specify a port range and common option together.")
			os.Exit(1)
		}
		if *portsPtr != "" {
			if strings.Contains(*portsPtr, "-") && strings.Contains(*portsPtr, ",") {
				fmt.Println("You can specify a ports range or an array, not both.")
				os.Exit(1)
			}
			if strings.Contains(*portsPtr, "-") {
				portsRange := string(*portsPtr)
				StartPort, EndPort = checkPortsRange(portsRange, StartPort, EndPort)
				portArrayBool = false
			} else if strings.Contains(*portsPtr, ",") {
				portsArray = checkPortsArray(*portsPtr)
				portArrayBool = true
			} else {
				portsRange := string(*portsPtr)
				StartPort, EndPort = checkPortsRange(portsRange, StartPort, EndPort)
				portArrayBool = false
			}
		}

		if !isURL(*portTargetPtr) {
			fmt.Println("The inputted target is not valid.")
			os.Exit(1)
		}
		if !outputFormatIsOk(*portOutputPtr) {
			fmt.Println("The output format is not valid.")
			os.Exit(1)
		}
	}

	if dirCommand.Parsed() {

		if *dirTargetPtr == "" {
			dirCommand.PrintDefaults()
			os.Exit(1)
		}
		if !isURL(*dirTargetPtr) {
			fmt.Println("The inputted target is not valid.")
			os.Exit(1)
		}
		if !outputFormatIsOk(*dirOutputPtr) {
			fmt.Println("The output format is not valid.")
			os.Exit(1)
		}
		if *dirIgnorePtr != "" {
			toBeIgnored := string(*dirIgnorePtr)
			dirIgnore = checkIgnore(toBeIgnored)
		}
	}

	if helpCommand.Parsed() {
		help()
		os.Exit(0)
	}

	if examplesCommand.Parsed() {
		examples()
		os.Exit(0)
	}

	result := Input{
		*reportTargetPtr,
		*reportWordlistDirPtr,
		*reportWordlistSubdomainPtr,
		*reportOutputPtr,
		reportIgnoreDir,
		reportIgnoreSub,
		*reportCrawlerDirPtr,
		*reportCrawlerSubdomainPtr,
		*reportSubdomainDBPtr,
		*reportCommonPtr,
		*dnsTargetPtr,
		*dnsOutputPtr,
		*dnsPlainPtr,
		*subdomainTargetPtr,
		*subdomainWordlistPtr,
		*subdomainOutputPtr,
		subdomainIgnore,
		*subdomainCrawlerPtr,
		*subdomainDBPtr,
		*subdomainPlainPtr,
		*dirTargetPtr,
		*dirWordlistPtr,
		*dirOutputPtr,
		dirIgnore,
		*dirCrawlerPtr,
		*dirPlainPtr,
		*portTargetPtr,
		*portOutputPtr,
		StartPort,
		EndPort,
		portArrayBool,
		portsArray,
		*portCommonPtr,
		*portPlainPtr,
	}
	return result
}

func checkIgnore(input string) []string {
	result := []string{}
	temp := strings.Split(input, ",")
	temp = removeDuplicateValues(temp)
	for _, elem := range temp {
		elem := strings.TrimSpace(elem)
		if len(elem) != 3 {
			fmt.Println("The status code you entered is invalid (It should consist of three characters).")
			os.Exit(1)
		}
		if ignoreInt, err := strconv.Atoi(elem); err == nil {
			if 100 <= ignoreInt && ignoreInt <= 599 {
				result = append(result, elem)
			} else {
				fmt.Println("The status code you entered is invalid (100 <= code <= 599).")
				os.Exit(1)
			}
		} else if strings.Contains(elem, "*") {
			if ignoreClassOk(elem) {
				result = append(result, elem)
			} else {
				fmt.Println("The status code you entered is invalid. You can enter * only as 1**,2**,3**,4**,5**.")
				os.Exit(1)
			}
		}
	}
	result = removeDuplicateValues(result)
	result = deleteUnusefulIgnoreresponses(result)
	return result
}

func deleteUnusefulIgnoreresponses(input []string) []string {
	var result []string
	toberemoved := []string{}
	classes := []string{}
	for _, elem := range input {
		if strings.Contains(elem, "*") {
			classes = append(classes, elem)
		}
	}

	for _, class := range classes {
		for _, elem := range input {
			if class[0] == elem[0] && elem[1] != '*' {
				toberemoved = append(toberemoved, elem)
			}
		}
	}

	result = Difference(input, toberemoved)
	return result
}

func Difference(a, b []string) (diff []string) {
	m := make(map[string]bool)
	for _, item := range b {
		m[item] = true
	}
	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}


func ignoreClassOk(input string) bool {
	if strings.Contains(input, "*") {
		if _, err := strconv.Atoi(string(input[0])); err == nil {
			i, err := strconv.Atoi(string(input[0]))
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			if i >= 1 && i <= 5 {
				if input[1] == byte('*') && input[2] == byte('*') {
					return true
				}
			}
		}
	}
	return false
}

func checkPortsArray(input string) []int {
	delimiter := byte(',')
	sliceOfPorts := strings.Split(input, string(delimiter))
	sliceOfPorts = removeDuplicateValues(sliceOfPorts)
	result := []int{}
	for _, elem := range sliceOfPorts {
		try, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Println("The inputted ports array is not valid.")
			os.Exit(1)
		}
		if try > 0 && try < 65536 {
			result = append(result, try)
		}
	}
	return result
}