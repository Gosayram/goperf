package request

import (
	"net/http"
	"time"
)

// IterateReqResp represents the performance metrics for a single URL across multiple requests
// It tracks response times, status codes, and byte counts for analysis
type IterateReqResp struct {
	URL         string          `json:"url"`
	Status      []int           `json:"status"`
	RespTimes   []time.Duration `json:"respTimes"`
	NumRequests int             `json:"numRequests"`
	Bytes       int             `json:"bytes"`
}

// IterateReqRespAll represents the complete performance test results including base URL and assets
// It combines metrics from the main page and all discovered assets (JS, CSS, images)
type IterateReqRespAll struct {
	AvgTotalRespTime       time.Duration    `json:"avgTotalRespTime"`
	AvgTotalLinearRespTime time.Duration    `json:"avgTotalLinearRespTime"`
	BaseURL                IterateReqResp   `json:"baseURL"`
	JSResps                []IterateReqResp `json:"jsResponses"`
	CSSResps               []IterateReqResp `json:"cssResponses"`
	IMGResps               []IterateReqResp `json:"imgResponses"`
}

// Result represents a single HTTP request result with detailed timing and response information.
// This structure is used to capture comprehensive metrics for performance analysis
// including URL and response status, timing information, response body and headers,
// and error details if any.
type Result struct {
	Total     time.Duration
	Average   time.Duration
	Channel   int
	Responses []*http.Response
}
