package request

import (
	"io"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"
)

/*
FetchInput used used as input to the Fetch method.

Structure Overview
  - BaseURL - the url to fetch
  - Retdat - if true then the document data is returned
  - Cookies - a cookie string to set on each request
  - UserAge - default is golang, but can be set to anything.`
*/
type FetchInput struct {
	BaseURL   string
	Retdat    bool
	Cookies   string
	Headers   string
	UserAgent string
}

/*
FetchResponse is the struct returned for Fetch()

Structure Overview:
  - Url - url of the fetched object <br>
  - Body - The body of the returned document.  Could be anything...html, js, css, etc.
  - Headers - The headers for the HttpResp object
  - Bytes - The number of bytes returned
  - Runes - The number of runes returned
  - Time - How long the Resp took.
  - Statue - the HttpResp status code.
  - Error - Any errors that were returned
*/
type FetchResponse struct {
	URL     string              `json:"url"`
	Resp    *http.Response      `json:"resp"`
	Body    string              `json:"body"`
	Headers map[string][]string `json:"headers"`
	Bytes   int                 `json:"bytes"`
	Runes   int                 `json:"runes"`
	Time    time.Duration       `json:"time"`
	Status  int                 `json:"status"`
	Error   string              `json:"error"`
}

/*
Fetch is a document from a url.  The document can be almost anything such as html, js, css, xml, etc.
A FetchInput object is used to encapsulate the various properties of the request.
*/
func Fetch(input FetchInput) *FetchResponse {
	url := input.BaseURL
	retdat := input.Retdat
	cookies := input.Cookies
	headers := strings.Split(input.Headers, "=")

	// Set up the http request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, http.NoBody)

	// Set the header only if we have a valid key=value format
	if len(headers) >= 2 && headers[0] != DefaultEmptyString {
		req.Header.Add(headers[0], headers[1])
	}

	// Set the use user-agent.  Default is 'goperf'.  So most like users will want to change it to something different.
	// Example user-agent: "Chrome/61.0.3163.100 Mobile Safari/537.36"
	req.Header.Add(UserAgentHeader, input.UserAgent)

	// cookies (example):  "sessionid_vagrant=5i4bgzvc0vy8xjgf1flfoh89cwsg74hz;
	// csrftoken_vagrant=taZjH9jskTjfbvDDq7OzdtQnTaB72zIk"
	req.Header.Add(CookieHeader, cookies)

	// Fetch the url and time the request
	start := time.Now()
	resp, err := client.Do(req)

	if err != nil {
		return &FetchResponse{
			URL:    err.Error(),
			Status: HTTPStatusConnectionError,
			Error:  ErrorRequestFailed,
		}
	}
	defer resp.Body.Close()
	responseTime := time.Since(start)

	// Read the html 'body' content from the response object
	body, err := io.ReadAll(resp.Body)
	Error := ""
	if err != nil {
		body = []byte("")
		Error = err.Error()
	}
	// This contains the text of the response.  HTML, Json, Exception, etc
	responseBody := string(body)

	output := FetchResponse{
		URL:     url,
		Resp:    resp,
		Body:    responseBody,
		Headers: resp.Header,
		Bytes:   len(responseBody),
		Runes:   utf8.RuneCountInString(responseBody),
		Time:    responseTime,
		Status:  resp.StatusCode,
		Error:   Error,
	}

	if !retdat { // we don't want the document data or headers
		output.Body = DefaultEmptyString
		output.Headers = make(map[string][]string)
	}
	// Close the response body and return the output
	return &output
}
