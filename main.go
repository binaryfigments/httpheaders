package httpheaders

import (
	"net/http"
	"strings"
)

// Data struct to fill with information.
type Data struct {
	Error        bool        `json:"error,omitempty"`
	ErrorMessage string      `json:"errormessage,omitempty"`
	Headers      http.Header `json:"headers,omitempty"`
	// TLS          *tls.ConnectionState `json:"tls,omitempty"`
}

// Get function that gets the Data
func Get(webappurl string) Data {

	// nextURL prefix check for incomplete
	if caseInsenstiveContains(webappurl, "http://") == false && caseInsenstiveContains(webappurl, "https://") == false {
		// TODO: Set warning
		webappurl = "http://" + webappurl
	}

	req, err := http.NewRequest("GET", webappurl, nil)
	req.Header.Add("User-Agent", "ocsr.nl checker")
	if err != nil {
		checkResult := Data{
			Error:        true,
			ErrorMessage: err.Error(),
		}
		return checkResult
	}
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		checkResult := Data{
			Error:        true,
			ErrorMessage: err.Error(),
		}
		return checkResult
	}
	resp.Body.Close()

	checkResult := Data{
		Error:   false,
		Headers: resp.Header,
		// TLS:     resp.TLS,
	}
	return checkResult
}

func caseInsenstiveContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}
