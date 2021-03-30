package crystals

import (
	"fmt"
	"io"
	"net/http"
)

func getPage(url string) (string, error) {
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if r.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http error: %v", http.StatusText)
	}

	defer r.Body.Close()
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func getList(url, start, end string) (string, error) {
	body, err := getPage(url)
	if err != nil {
		return "", err
	}


	for

}
