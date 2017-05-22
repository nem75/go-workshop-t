package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/prometheus/common/log"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.Method)
		log.Info(r.FormValue("key"))
	}))
	defer ts.Close()

	// res, err := http.Get(ts.URL)

	form := url.Values{}
	form.Set("key", "value1")

	req, _ := http.NewRequest("POST", ts.URL, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	hc := http.Client{}
	log.Info("form was %v", form)
	hc.Do(req)

	// fmt.Println(resp.StatusCode, resp.Header)
}
