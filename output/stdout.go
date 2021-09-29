package output

import (
	"fmt"

	"github.com/mgutz/ansi"
)

func GetPrinter(name string) Printer {
	if output, ok := printers[name]; ok {
		return output
	} else {
		panic("No output by name " + name)
	}
}

type StdOut struct {
	ansiGreen func(string) string
	ansiRed   func(string) string
}

func (stdout StdOut) Print(str string) {
	fmt.Println(str)
}

func (stdout StdOut) PrintVulnFound(vuln string) {
	fmt.Println(stdout.ansiGreen("Vulnerable") + " : " + vuln)
}

func (stdout StdOut) PrintVulnNotFound(vuln string) {
	fmt.Println(stdout.ansiRed("Not vulnerable") + " : " + vuln)
}

var StdOutPrinter = StdOut{
	ansiGreen: ansi.ColorFunc("green+h"),
	ansiRed:   ansi.ColorFunc("red+h"),
}
