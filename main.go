package httpheaders

import (
	"net/http"
	"strings"
)

// ResponseHeaders struct to fill with information.
type ResponseHeaders struct {
	Error        bool        `json:"error,omitempty"`
	ErrorMessage string      `json:"errormessage,omitempty"`
	Headers      http.Header `json:"headers,omitempty"`
}

// Get function that gets the ResponseHeaders
func Get(webappurl string) ResponseHeaders {

	// nextURL prefix check for incomplete
	if caseInsenstiveContains(webappurl, "http://") == false && caseInsenstiveContains(webappurl, "https://") == false {
		// TODO: Set warning
		webappurl = "http://" + webappurl
	}

	req, err := http.NewRequest("GET", webappurl, nil)
	req.Header.Add("User-Agent", "ocsr.nl checker")
	if err != nil {
		checkResult := ResponseHeaders{
			Error:        true,
			ErrorMessage: err.Error(),
		}
		return checkResult
	}
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		checkResult := ResponseHeaders{
			Error:        true,
			ErrorMessage: err.Error(),
		}
		return checkResult
	}
	resp.Body.Close()

	checkResult := ResponseHeaders{
		Error:   false,
		Headers: resp.Header,
	}
	return checkResult
}

func caseInsenstiveContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}
