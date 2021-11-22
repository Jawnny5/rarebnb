package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key string
	value string 
}

var theTests = []struct {
	name string
	url string
	method string
	params []postData
	expectedStatusCode int 
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"hurtyurt", "/hurtyurt", "GET", []postData{}, http.StatusOK},
	{"timothys", "/timothys", "GET", []postData{}, http.StatusOK},
	{"searchavailability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"makereservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-search-avail", "/search-availability", "POST", []postData{
		{key: "start", value: "01-01-2020"},
		{key: "end", value: "01-03-2020"},
	}, http.StatusOK},
	{"post-search-avail-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "01-01-2020"},
		{key: "end", value: "01-03-2020"},
	}, http.StatusOK},
	{"make-reservation-post", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Sir"},
		{key: "last_name", value: "SAL"},
		{key: "email", value: "ssmal@niceguy.com"},
		{key: "phone", value: "1-800-NICE-GUY"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T){
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("For %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
				values := url.Values{}
				for _, x := range e.params{
					values.Add(x.key, x.value)
				}
				resp, err := ts.Client().PostForm(ts.URL + e.url, values)
				if err != nil {
					t.Log(err)
					t.Fatal(err)
				}
				if resp.StatusCode != e.expectedStatusCode {
					t.Errorf("For %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
				}
		}
	}
}