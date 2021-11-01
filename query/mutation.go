package query

import (
	"context"
	"fmt"
	"juuri/options"
	"strings"

	"github.com/machinebox/graphql"
)

type MutationQuery struct {
	query string
}

var MutationCheck = MutationQuery{
	query: `
	{
		__schema {
		  mutationType {
			fields {
			  name
			}
		  }
		}
	}
	`,
}

type MutationQueryResponse struct {
	Schema struct {
		MutationType *struct {
			Fields []struct {
				Name string `json:"name"`
			} `json:"fields"`
		} `json:"mutationType"`
	} `json:"__schema"`
}

func (q MutationQuery) Check(client *graphql.Client, options options.JuuriOptions) (bool, string) {
	var resp MutationQueryResponse
	req := graphql.NewRequest(q.query)
	setGraphQLRequestHeaders(req, options.Headers, options.Debug)
	if err := client.Run(context.Background(), req, &resp); err != nil {
		if options.Debug {
			fmt.Printf("Error in %s: %s", q.Describe(), err.Error())
		}
	}

	var b strings.Builder
	for _, s := range resp.Schema.MutationType.Fields {
		b.WriteString(s.Name + "\n")
	}

	return resp.Schema.MutationType != nil, b.String()
}

func (q MutationQuery) Describe() string {
	return "Mutations"
}
