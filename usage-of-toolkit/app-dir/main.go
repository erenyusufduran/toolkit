package main

import "github.com/erenyusufduran/toolkit"

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExists("./test-dir")
}
