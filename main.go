package main

import (
	"flag"
	"fmt"
	"juuri/options"
	"juuri/output"
	"juuri/query"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const BANNER = `
      _               _ 
     (_)_ ____ ______(_)
    / / // / // / __/ /
 __/ /\_,_/\_,_/_/ /_/
|___/
`

const VOYAGER_URL = "http://apis.guru/graphql-voyager/?url="

func usage() {
	fmt.Println("juuri OPTIONS <url>")
	flag.PrintDefaults()
}

func printBanner() {
	fmt.Println(BANNER)
}

func parseHeaders(headersStr string) map[string]string {
	headerMap := map[string]string{}
	headers := strings.Split(headersStr, ",")
	for _, header := range headers {
		headerNameValue := strings.Split(header, ":")
		headerMap[headerNameValue[0]] = headerNameValue[1]

	}
	return headerMap
}

func main() {
	printBanner()

	var options = options.JuuriOptions{}
	var headers string
	flag.BoolVar(&options.Debug, "debug", false, "Debug logging")
	flag.BoolVar(&options.OpenIntrospectionInVoyager, "open-in-voyager", false, "Open introspection result in GraphQL Voyager")
	flag.StringVar(&options.File, "file", "", "Output file")
	flag.StringVar(&headers, "headers", "", "List of HTTP headers separated by comma, e.g. headers=accept-encoding:gzip,content-type:application/json")
	flag.Usage = usage
	flag.Parse()
	if len(headers) > 0 {
		options.Headers = parseHeaders(headers)
	}
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	urlArg := flag.Arg(0)
	_, err := url.ParseRequestURI(urlArg)
	if err != nil {
		panic("Invalid URL " + urlArg)
	}

	var printer output.Printer
	if len(options.File) > 0 {
		printer = output.FileOutPrinter
	} else {
		printer = output.StdOutPrinter
	}
	printer.Init(&options)

	for _, check := range query.VulnChecks {
		vulnerable, text := check.Check(urlArg, options)
		if vulnerable {
			printer.PrintVulnFound(check.Describe())
			if len(text) > 0 {
				printer.Print(text)
			}
		} else {
			printer.PrintVulnNotFound(check.Describe())
		}
	}
	printer.Stop()

	if options.OpenIntrospectionInVoyager {
		fullVoyagerUrl := VOYAGER_URL + urlArg

		fmt.Println("Opening API in GraphQL Voyager")
		var browserErr error
		switch runtime.GOOS {
		case "linux":
			browserErr = exec.Command("xdg-open", fullVoyagerUrl).Start()
		case "windows":
			browserErr = exec.Command("rundll32", "url.dll,FileProtocolHandler", fullVoyagerUrl).Start()
		case "darwin":
			browserErr = exec.Command("open", fullVoyagerUrl).Start()
		default:
			browserErr = fmt.Errorf("Unsupported platform")
		}
		if browserErr != nil {
			fmt.Printf("Error opening browser for GraphQL Voyager: %s\n", browserErr.Error())
		}
	}
}
