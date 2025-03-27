package main

import (
	"testing"
)

func TestScraperMap(t *testing.T) {
	for k := range ScraperMap {
		_, exists := ScraperMap[k]
		if !exists {
			t.Errorf("expected scraper for key %s, but got none", k)

		}
	}
}
