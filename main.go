package main

import (
	"flag"
	"fmt"
	"juuri/options"
	"juuri/output"
	"juuri/query"
	"os"

	"github.com/machinebox/graphql"
)

const BANNER = `
      _               _ 
     (_)_ ____ ______(_)
    / / // / // / __/ /
 __/ /\_,_/\_,_/_/ /_/
|___/
`

func usage() {
	fmt.Println("juuri OPTIONS <url>")
	flag.PrintDefaults()
}

func printBanner() {
	fmt.Println(BANNER)
}

func main() {
	printBanner()

	var options = options.JuuriOptions{}
	var printerType string
	flag.BoolVar(&options.Debug, "debug", false, "Debug logging")
	flag.StringVar(&printerType, "output", "stdout", "Output type: currently only \"stdout\"")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	client := graphql.NewClient(flag.Arg(0))

	printer := output.GetPrinter(printerType)

	for _, check := range query.VulnChecks {
		vulnerable, text := check.Check(client, options)
		if vulnerable {
			printer.PrintVulnFound(check.Describe())
			if len(text) > 0 {
				printer.Print(text)
			}
		} else {
			printer.PrintVulnNotFound(check.Describe())
		}
	}
}
