package main

const (
	sourceSite = `https://thecrystalcouncil.com/crystals`
	tagStart   = `<h5 class="ml-2 mt-2 mb-0 text-dark">`
	tagEnd     = `</h5>`
)

func main() {
	crystals.getList(sourceSite)
}
