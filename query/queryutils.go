package query

import (
	"fmt"

	"github.com/machinebox/graphql"
)

func setGraphQLRequestHeaders(req *graphql.Request, headers map[string]string, debug bool) {
	if len(headers) > 0 {
		for headerName, headerValue := range headers {
			if debug {
				fmt.Printf("Setting request header %s to %s\n", headerName, headerValue)
			}
			req.Header.Add(headerName, headerValue)
		}
	}
}
