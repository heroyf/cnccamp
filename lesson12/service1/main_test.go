package main

import (
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	// test header
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/header", nil)
	req.Header.Set("test-header", "test")
	resp, _ := (&http.Client{}).Do(req)
	if resp.Header.Get("test-header") != "[test]" {
		t.Errorf("get test-header failed: %v", resp.Header.Get("test-header"))
	} else {

	}

	if _, ok := resp.Header["Sys-Version"]; !ok {
		t.Errorf("get version-header failed")
	}
}
