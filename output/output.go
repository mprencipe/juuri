package output

type Output struct{}

type Printer interface {
	Print(str string)
	PrintVulnFound(vuln string)
	PrintVulnNotFound(vuln string)
}

var printers = map[string]Printer{
	"stdout": StdOutPrinter,
}
