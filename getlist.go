// Package crystals is a sample implementation of web data scraping
package crystals

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetPage(url string) (string, error) {
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if r.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http error: %s", http.StatusText)
	}

	defer r.Body.Close()
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func GetList(url, start, end string) ([]string, error) {
	list := []string{}
	contents, err := GetPage(url)
	if err != nil {
		return nil, err
	}

	var a, b int

	for {
		a = strings.Index(contents, start) + 1
		if a < 1 {
			break
		}
		b = strings.Index(contents[a:], end) - 1
		if b < 1 {
			b = len(contents) - 1
		}

		item := contents[a:b]

		list = append(list, item)

	}
	return list, nil
}
