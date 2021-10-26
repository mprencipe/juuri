package output

import (
	"fmt"
	"juuri/options"
	"os"
)

type FileOut struct {
	file *os.File
}

func (out *FileOut) Init(options *options.JuuriOptions) {
	fmt.Printf("Writing output to %s\n", options.File)
	file, err := os.Create(options.File)
	if err != nil {
		panic(err)
	}
	out.file = file
}

func (out FileOut) Print(str string) {
	_, err := out.file.WriteString(fmt.Sprintf("%s\n", str))
	if err != nil {
		panic(err)
	}
}

func (out FileOut) PrintVulnFound(vuln string) {
	_, err := out.file.WriteString(fmt.Sprintf("Vulnerable: %s\n", vuln))
	if err != nil {
		panic(err)
	}
}

func (out FileOut) PrintVulnNotFound(vuln string) {
	_, err := out.file.WriteString(fmt.Sprintf("Not vulnerable: %s\n", vuln))
	if err != nil {
		panic(err)
	}
}

func (out FileOut) Stop() {
	defer out.file.Close()
	err := out.file.Sync()
	if err != nil {
		panic(err)
	}
}

var FileOutPrinter = &FileOut{}
