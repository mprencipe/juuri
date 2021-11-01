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

func (q MutationQuery) Check(url string, options options.JuuriOptions) (bool, string) {
	client := graphql.NewClient(url)
	var resp MutationQueryResponse
	req := graphql.NewRequest(q.query)
	setGraphQLRequestHeaders(req, options.Headers, options.Debug)
	if err := client.Run(context.Background(), req, &resp); err != nil {
		if options.Debug {
			fmt.Printf("Error in %s: %s", q.Describe(), err.Error())
		}
	}

	if resp == (MutationQueryResponse{}) {
		return false, ""
	}

	mutations := make([]string, 0)
	for _, field := range resp.Schema.MutationType.Fields {
		mutations = append(mutations, field.Name)
	}
	return true, strings.Join(mutations, ", ")
}

func (q MutationQuery) Describe() string {
	return "Mutations"
}
