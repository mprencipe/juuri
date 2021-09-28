package main

import (
	"flag"
	"fmt"
	"juuri/options"
	"juuri/query"
	"os"

	"github.com/machinebox/graphql"
	"github.com/mgutz/ansi"
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
	var options = options.JuuriOptions{}
	flag.BoolVar(&options.Debug, "debug", false, "Debug logging")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	printBanner()

	client := graphql.NewClient(flag.Arg(0))

	ansiGreen := ansi.ColorFunc("green+h")
	ansiRed := ansi.ColorFunc("red+h")

	for _, check := range query.VulnChecks {
		result := check.Check(client, options)
		var resultAnsi string
		if result {
			resultAnsi = ansiGreen("VULNERABLE")
		} else {
			resultAnsi = ansiRed("NOT VULNERABLE")
		}
		fmt.Printf("%s %s\n", resultAnsi, check.Describe())
	}
}
