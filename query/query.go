package query

import (
	"juuri/options"

	"github.com/machinebox/graphql"
)

type VulnCheck interface {
	Check(client *graphql.Client, options options.JuuriOptions) bool
	Describe() string
}

var VulnChecks = []VulnCheck{
	MutationCheck,
	IntrospectionCheck,
}
