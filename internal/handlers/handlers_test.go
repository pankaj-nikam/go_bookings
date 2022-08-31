package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},

	{"post-search-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-01"},
	}, http.StatusOK},

	{"post-search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-01"},
	}, http.StatusOK},

	{"make-reservation-post", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Pankaj"},
		{key: "last_name", value: "Nikam"},
		{key: "email", value: "demo@demo.com"},
		{key: "phone", value: "555-5555"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range theTests {
		if test.method == "GET" {
			response, err := ts.Client().Get(ts.URL + test.url)
			if err != nil {
				t.Error("error occurred")
			} else {
				if response.StatusCode != test.expectedStatusCode {
					t.Errorf("expected %d and got %d", test.expectedStatusCode, response.StatusCode)
				}
			}
		} else if test.method == "POST" {
			values := url.Values{}
			for _, x := range test.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+test.url, values)
			if err != nil {
				t.Error("error occurred")
			} else {
				if resp.StatusCode != test.expectedStatusCode {
					t.Errorf("expected %d and got %d", test.expectedStatusCode, resp.StatusCode)
				}
			}
		}
	}
}
