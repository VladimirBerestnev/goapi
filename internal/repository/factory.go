package repository

import (
	"context"
	"go/rest/internal/app/files"
	"go/rest/internal/entity"
	"go/rest/internal/repository/postgres"
	"go/rest/internal/repository/sqlite"
	_ "go/rest/migrations"

	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

type IDatabase interface {
	Create(context.Context, entity.Task) error
	Get(context.Context) ([]entity.Task, error)
	Delete(context.Context, string) error
	Update(context.Context, entity.Task) error
}

func New() IDatabase {
	data := files.OpenYaml()
	dbType, exist := data["dbType"]
	if !exist {
		panic("dbType not exist")
	}

	switch dbType {
	case "sqlite":
		return sqlite.New()
	case "postgres":
		return postgres.New()
	default:
		panic("dbType incorrect")
	}

}
