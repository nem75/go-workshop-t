package main

import (
	"fmt"
	st "go-workshop-t/hw_4/store"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_handleRequestPost(t *testing.T) {
	// Prevent file system access.
	kvs = st.NewStore(st.NewVirtualRespository())

	// Start new server with handler to be tested.
	ts := httptest.NewServer(http.HandlerFunc(handlePost))
	defer ts.Close()

	// Build test request.
	form := url.Values{}
	form.Set("key", "value1")
	req, _ := http.NewRequest("POST", ts.URL, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send request.
	hc := http.Client{}
	resp, _ := hc.Do(req)

	// Test result code.
	ex := 201
	if rs := resp.StatusCode; rs != ex {
		printMsg(ex, rs)
		t.Fail()
	}
}

func printMsg(ex interface{}, rs interface{}) {
	fmt.Printf("Result was '%v' (%T) instead of '%v' (%T)\n", rs, rs, ex, ex)
}
