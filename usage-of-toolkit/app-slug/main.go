package main

import (
	"log"

	"github.com/erenyusufduran/toolkit"
)

func main() {
	toSlug := "NOW!!_ is the time 123"
	var tools toolkit.Tools

	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}
	log.Println(slugified)
}
