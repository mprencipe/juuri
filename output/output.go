package output

import (
	"juuri/options"
)

type Output struct{}

type Printer interface {
	Init(options *options.JuuriOptions)
	Print(str string)
	PrintVulnFound(vuln string)
	PrintVulnNotFound(vuln string)
	Stop()
}
