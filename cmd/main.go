package main

import (
	"go/rest/internal/app/files"
	"go/rest/internal/app/routing"
)

func main() {

	files.LoadFromFile()
	routing.APIrout()

}
