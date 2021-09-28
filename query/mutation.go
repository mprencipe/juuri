package query

import (
	"context"
	"fmt"
	"juuri/options"

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

func (q MutationQuery) Check(client *graphql.Client, options options.JuuriOptions) bool {
	var resp MutationQueryResponse
	req := graphql.NewRequest(q.query)

	if err := client.Run(context.Background(), req, &resp); err != nil {
		if options.Debug {
			fmt.Printf("Error in %s: %s", q.Describe(), err.Error())
		}
	}

	return resp.Schema.MutationType != nil
}

func (q MutationQuery) Describe() string {
	return "Mutations"
}
