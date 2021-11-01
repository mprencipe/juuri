package query

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"juuri/options"
	"net/http"
)

type BatchingQuery struct {
	query string
}

var BatchingCheck = BatchingQuery{
	query: `{"query": "query { test: Query { foo } test2: Query { bar } }","variables": null}`,
}

type BatchingQueryResponse struct {
	Errors []struct {
		Message string
	} `json:"errors"`
}

func createRequest(url string, queryStr string) *http.Request {
	var query = []byte(queryStr)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(query))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	return req
}

func readJson(respBody io.ReadCloser, resp *BatchingQueryResponse) {
	body, ioErr := ioutil.ReadAll(respBody)
	if ioErr != nil {
		panic(ioErr)
	}
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		panic(jsonErr)
	}
}

func (q BatchingQuery) Check(url string, options options.JuuriOptions) (bool, string) {
	var jsonResp BatchingQueryResponse
	req := createRequest(url, q.query)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	readJson(resp.Body, &jsonResp)
	return len(jsonResp.Errors) == 2, ""
}

func (q BatchingQuery) Describe() string {
	return "Batching"
}
