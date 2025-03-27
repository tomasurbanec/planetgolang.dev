package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScrape(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") != userAgent {
			t.Errorf("Expected User-Agent: %s, got: %s", userAgent, r.Header.Get("User-Agent"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"planet":"golang"}`))
	}))
	defer server.Close()

	body, err := Scrape(server.URL)

	if err != nil {
		t.Errorf("Got an error scraping: %s", err)
	}

	if body == "" {
		t.Error("Got empty body scraping")
	}
}
