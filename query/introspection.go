package query

import (
	"context"
	"fmt"
	"juuri/options"

	"github.com/machinebox/graphql"
)

type IntrospectionQuery struct {
	query string
}

var IntrospectionCheck = IntrospectionQuery{
	query: `
	{
		__schema {
			queryType {name}
			subscriptionType{name}
			types{...FullType}
			directives{
				name
				description
				locations
				args{...InputValue}
			}
		}
	}
	fragment FullType on __Type{
		kind
		name
		description
		fields(includeDeprecated:true){
			name
			description
			args{...InputValue}
			type{...TypeRef}isDeprecated deprecationReason}
			inputFields{...InputValue}
			interfaces{...TypeRef}
			enumValues(includeDeprecated:true){
				name
				description
				isDeprecated
				deprecationReason
			}
			possibleTypes{...TypeRef}
		}
		fragment InputValue on __InputValue{
			name
			description
			type{...TypeRef}defaultValue
		}
		fragment TypeRef on __Type{
			kind
			name
			ofType{
				kind
				name
				ofType{
					kind
					name
					ofType{
						kind
						name
						ofType{
							kind
							name
							ofType{
								kind
								name
								ofType{
									kind
									name
									ofType{
										kind
										name
									}
								}
							}
						}
					}
				}
			}
		}
	`,
}

type IntrospectionQueryResponse struct {
	Schema struct {
		Types []struct {
			Name string `json:"name"`
		} `json:"types"`
	} `json:"__schema"`
}

func (q IntrospectionQuery) Check(url string, options options.JuuriOptions) (bool, string) {
	client := graphql.NewClient(url)
	var resp IntrospectionQueryResponse
	req := graphql.NewRequest(q.query)
	setGraphQLRequestHeaders(req, options.Headers, options.Debug)
	if err := client.Run(context.Background(), req, &resp); err != nil {
		if options.Debug {
			fmt.Printf("Error in %s: %s\n", q.Describe(), err.Error())
		}
		return false, ""
	}

	return resp.Schema.Types != nil, ""
}

func (q IntrospectionQuery) Describe() string {
	return "Introspection"
}
