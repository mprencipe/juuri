package output

import (
	"fmt"

	"juuri/options"

	"github.com/mgutz/ansi"
)

type StdOut struct {
	ansiGreen func(string) string
	ansiRed   func(string) string
}

func (out StdOut) Init(options *options.JuuriOptions) {}

func (out StdOut) Print(str string) {
	fmt.Println(str)
}

func (out StdOut) PrintVulnFound(vuln string) {
	fmt.Println(out.ansiGreen("Vulnerable") + " : " + vuln)
}

func (out StdOut) PrintVulnNotFound(vuln string) {
	fmt.Println(out.ansiRed("Not vulnerable") + " : " + vuln)
}

func (out StdOut) Stop() {}

var StdOutPrinter = StdOut{
	ansiGreen: ansi.ColorFunc("green+h"),
	ansiRed:   ansi.ColorFunc("red+h"),
}
