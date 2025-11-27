package handlers

import (
	"net/http"
)

const (
	available    = "available"
	notAvailable = "not available"
)

func checkAvailability(links []string) (map[string]string, error) {
	checkedLinks := make(map[string]string)

	for _, link := range links {
		client := &http.Client{}

		resp, err := client.Get("https://" + link)
		if err != nil {
			resp, err = client.Get("http://" + link)
			if err != nil {
				checkedLinks[link] = notAvailable
				continue
			}

		}
		defer resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode <= 300 {
			checkedLinks[link] = available
		} else {
			checkedLinks[link] = notAvailable
		}

	}

	return checkedLinks, nil
}
