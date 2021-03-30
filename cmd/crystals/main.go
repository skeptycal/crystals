package main

import (
	"fmt"
	"strings"

	"github.com/skeptycal/crystals"
	"github.com/skeptycal/util/ansi"
)

const (
	sourceSite = `https://thecrystalcouncil.com/crystals`
	tagStart   = `<h5 class="ml-2 mt-2 mb-0 text-dark">`
	tagEnd     = `</h5>`
)

func AlternateColors(s string) string {
	for _, line := range strings.Split(s, "\n") {
		ansi.Print(color, line)
	}
}

func main() {

	fmt.Println("site: ", sourceSite)
	fmt.Println("tagStart: ", tagStart)
	fmt.Println("tagEnd: ", tagEnd)

	page, err := crystals.GetPage(sourceSite)
	if err != nil {
		fmt.Println(fmt.Errorf("error retrieving page contents: %v", err))
	}

	fmt.Println()
	fmt.Println(page)

	// list, err := crystals.GetList(sourceSite, tagStart, tagEnd)
	// if err != nil {
	// 	os.Exit(1)
	// }
	// for _, item := range list {
	// 	fmt.Println(item)
	// }
}
