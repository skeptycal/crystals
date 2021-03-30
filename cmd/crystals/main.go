package main

import (
	"fmt"
	"os"

	"github.com/skeptycal/crystals"
)

const (
	sourceSite = `https://thecrystalcouncil.com/crystals`
	tagStart   = `<h5 class="ml-2 mt-2 mb-0 text-dark">`
	tagEnd     = `</h5>`
)

func main() {
	list, err := crystals.GetList(sourceSite, tagStart, tagEnd)
	if err != nil {
		os.Exit(1)
	}
	for _, item := range list {
		fmt.Println(item)
	}
}
