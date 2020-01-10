// +build integration

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testmain(t *testing.T) {
	fmt.Println("Hi There! I'm a Test.")
}

func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"status\": \"good\"}")
	}

	req := httptest.NewRequest("GET", "http://localhot:8080", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if 200 != resp.StatusCode {
		t.Fatal("Status Code Not Ok")
	}

}

func TestGet(t *testing.T) {
	resp := Get(server.URL + "/user")
	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}
