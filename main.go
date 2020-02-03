package httpheaders

import (
	"net/http"
	"strings"
	"time"
)

// Data struct to fill with information.
type Data struct {
	Status          int         `json:"status,omitempty"`
	Protocol        string      `json:"ptotocol,omitempty"`
	Error           bool        `json:"error,omitempty"`
	ErrorMessage    string      `json:"errormessage,omitempty"`
	Headers         http.Header `json:"headers,omitempty"`
	FollowRedirects bool
	// TLS          *tls.ConnectionState `json:"tls,omitempty"`
}

// Get function that gets the Data
func Get(webappurl string, followRedirects bool) Data {
	// nextURL prefix check for incomplete
	if caseInsenstiveContains(webappurl, "http://") == false && caseInsenstiveContains(webappurl, "https://") == false {
		// TODO: Set warning
		webappurl = "http://" + webappurl
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	if followRedirects == false {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	response, err := client.Get(webappurl)
	if err != nil {
		checkResult := Data{
			Error:        true,
			ErrorMessage: err.Error(),
		}
		return checkResult
	}
	response.Body.Close()

	checkResult := Data{
		Error:           false,
		FollowRedirects: followRedirects,
		Headers:         response.Header,
		Status:          response.StatusCode,
		Protocol:        response.Proto,
		// TLS:     resp.TLS,
	}
	return checkResult
}

func caseInsenstiveContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}
