package request

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gnulnx/color"

	"github.com/Gosayram/goperf/httputils"
)

// FetchAllResponse is the return structure from FetchAll
type FetchAllResponse struct {
	BaseURL         *FetchResponse  `json:"BaseURL"`
	Time            time.Duration   `json:"time"`
	TotalTime       time.Duration   `json:"totalTime"`
	TotalLinearTime time.Duration   `json:"totalLinearTime"`
	TotalBytes      int             `json:"totalBytes"`
	JSResponses     []FetchResponse `json:"jsResponses"`
	IMGResponses    []FetchResponse `json:"imgResponses"`
	CSSResponses    []FetchResponse `json:"cssResponses"`

	Body string `json:"body"`
}

/*
FetchAll takes a FetchInput object and proceeds to fetch
the BaseURL and then fetch all of it's assets.

Assets currently refer to script, style, and img tags.

Each asset class is fetched in it's own go routine.
If retdata is False we don't return the Body or Header.
This is useful if you only want the timing data.
For instance you might find it useful to fetch with retdat=true
the first time around to get all the data and write to file.
The subsequet requests could be used as part of a perf test where
you only need the raw timing and size data.  In those cases
you can set retdat=false to effectively cut down on the verbosity
*/
func FetchAll(input FetchInput) *FetchAllResponse {
	retdat := input.Retdat

	// Fetch initial url
	start := time.Now()
	input.Retdat = true
	output := Fetch(input)

	// Now parse output for js, css, img urls
	jsfiles, imgfiles, cssfiles := httputils.GetAssets(output.Body)

	// Now lets create some go routines and fetch all the js, img, css files
	c1 := make(chan []FetchResponse)
	c2 := make(chan []FetchResponse)
	c3 := make(chan []FetchResponse)

	go GoFetchAllAssetArray(jsfiles, input, c1)
	go GoFetchAllAssetArray(imgfiles, input, c2)
	go GoFetchAllAssetArray(cssfiles, input, c3)

	jsResponses := []FetchResponse{}
	imgResponses := []FetchResponse{}
	cssResponses := []FetchResponse{}

	for i := 0; i < AssetTypesCount; i++ {
		select {
		case jsResponses = <-c1:
		case imgResponses = <-c2:
		case cssResponses = <-c3:
		}
	}
	totalTime2 := time.Since(start)

	if !retdat {
		output.Body = DefaultEmptyString
		output.Headers = make(map[string][]string)
	}

	calcTotal := func(resp []FetchResponse) (time.Duration, int) {
		totalTime := time.Duration(0)
		totalBytes := 0
		for _, val := range resp {
			totalTime += val.Time
			totalBytes += val.Bytes
		}
		return totalTime, totalBytes
	}

	jsTime, jsBytes := calcTotal(jsResponses)
	cssTime, cssBytes := calcTotal(cssResponses)
	imgTime, imgBytes := calcTotal(imgResponses)

	totalLinearTime := output.Time + jsTime + cssTime + imgTime
	totalBytes := output.Bytes + jsBytes + cssBytes + imgBytes

	resp := FetchAllResponse{
		BaseURL:         output,
		Time:            output.Time,
		TotalTime:       totalTime2,
		TotalLinearTime: totalLinearTime,
		TotalBytes:      totalBytes,
		Body:            output.Body,
		JSResponses:     jsResponses,
		IMGResponses:    imgResponses,
		CSSResponses:    cssResponses,
	}

	return &resp
}

/*
PrintFetchAllResponse takes a FetchAllResponse object and prints the results to stdout
It displays detailed performance metrics with color-coded output for easy reading
*/
func PrintFetchAllResponse(resp *FetchAllResponse) {
	yel := color.New(color.FgHiYellow).SprintfFunc()
	yellow := color.New(color.FgHiYellow, color.Underline).SprintfFunc()
	grey := color.New(color.FgHiBlack).SprintfFunc()
	white := color.New(color.FgWhite).SprintfFunc()
	red := color.New(color.FgRed).SprintfFunc()
	green := color.New(color.FgGreen).SprintfFunc()
	total := resp.BaseURL.Time

	// Show the resp body.
	fmt.Print(resp.Body)

	color.Red("Base Url Results")
	if resp.BaseURL.Status == HTTPStatusOK {
		fmt.Printf(" - %-34s %-25s\n", yel("Status:"), green(strconv.Itoa(resp.BaseURL.Status)))
	} else {
		fmt.Printf(" - %-34s %-25s\n", yel("Status:"), red(strconv.Itoa(resp.BaseURL.Status)))
	}
	fmt.Printf(" - %-34s %-25s\n", yel("Url:"), white(resp.BaseURL.URL))
	fmt.Printf(" - %-34s %-25s\n", yel("Time to first byte"), total.String())
	fmt.Printf(" - %-34s %-25s\n", yel("Bytes"), strconv.Itoa(resp.BaseURL.Bytes))
	fmt.Printf(" - %-34s %-25s\n", yel("Runes"), strconv.Itoa(resp.BaseURL.Runes))
	fmt.Printf(" - %-34s %-25s\n", yel("TotalTime"), resp.TotalTime.String())
	fmt.Printf(" - %-34s %-25s\n", yel("TotalBytes"), strconv.Itoa(resp.TotalBytes))

	printAssets := func(title string, results []FetchResponse) {
		color.Red(title)
		fmt.Printf(" - %-24s %-22s %-21s\n", yellow("Time"), yellow("Bytes"), yellow("Url"))
		for i, val := range results {
			if i%2 == 0 {
				fmt.Printf(" - %-22s %-20s %-10s \n", white(val.Time.String()), white(strconv.Itoa(val.Bytes)), white(val.URL))
			} else {
				fmt.Printf(" - %-22s %-20s %-10s \n", grey(val.Time.String()), grey(strconv.Itoa(val.Bytes)), grey(val.URL))
			}
		}
	}

	printAssets("JS Responses", resp.JSResponses)
	printAssets("CSS Responses", resp.CSSResponses)
	printAssets("IMG Responses", resp.IMGResponses)
}

// DefineAssetURL is used to determine if an asset is a local resource
func DefineAssetURL(baseURL, asseturl string) string {
	/*
		If the url starts with a / we know it's a local resource
		so we prepend the baseURL to it
	*/
	if asseturl[:HTTPProtocolLength] == HTTPScheme {
		return asseturl
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	baseURL = fmt.Sprintf("%s://%s", u.Scheme, u.Hostname())

	if asseturl[FirstPathElement] == '/' {
		asseturl = baseURL + asseturl
	} else {
		asseturl = baseURL + PathSeparator + asseturl
	}
	return asseturl
}

// GoFetchAllAssetArray fetches all assets from the provided URLs concurrently
// It creates goroutines for each asset URL and returns the results via a channel
func GoFetchAllAssetArray(files []string, input FetchInput, resp chan []FetchResponse) {
	baseURL := input.BaseURL

	chanHolder := []chan FetchResponse{}
	for i, assetURL := range files {
		chanHolder = append(chanHolder, make(chan FetchResponse))
		go func(c chan FetchResponse, assetURL string, input FetchInput) {
			input.BaseURL = DefineAssetURL(baseURL, assetURL)
			c <- *Fetch(input)
		}(chanHolder[i], assetURL, input)
	}

	// Wait on all the channels
	responses := []FetchResponse{}
	for _, ch := range chanHolder {
		responses = append(responses, <-ch)
	}

	resp <- responses
}
