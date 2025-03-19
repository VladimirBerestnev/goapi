package main

import (
	"go/rest/internal/app/delivery"
	"go/rest/internal/app/routing"
	"go/rest/internal/repository"
	"go/rest/internal/usecase"
)

func main() {
	repo := repository.New()
	uc := usecase.New(repo)
	handler := delivery.New(uc)
	routing.APIrout(handler)
}
