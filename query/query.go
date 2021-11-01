package query

import (
	"juuri/options"
)

type VulnCheck interface {
	Check(url string, options options.JuuriOptions) (bool, string)
	Describe() string
}

var VulnChecks = []VulnCheck{
	MutationCheck,
	IntrospectionCheck,
	BatchingCheck,
}
