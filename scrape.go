package planetgolang

import (
	"io/ioutil"
	"net/http"
)

func Scrape(url string) (string, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("User-Agent", "planetgolang.dev/1.0")

	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
