package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"go/rest/internal/app/files"
	"go/rest/internal/entity"
	_ "go/rest/migrations"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

type Repo struct {
	db *sql.DB
}

func New() *Repo {
	data := files.OpenYaml()
	db, err := sql.Open("sqlite", data["db_path"])
	if err != nil {
		panic(err)
	}

	if err = UpMigrations(db); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return &Repo{db: db}
}

func (r *Repo) Create(c context.Context, task entity.Task) error {
	query := `INSERT INTO tasks (ID, TITLE, DESCRIPTION, STATUS, PRIORITY) VALUES (?, ?, ?, ?, ?);`

	_, err := r.db.ExecContext(c, query, task.ID, task.Title, task.Desc, task.Status, task.Priority)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) Update(c context.Context, task entity.Task) error {
	_, err := r.db.ExecContext(c, "UPDATE tasks SET TITLE=?, DESCRIPTION=?, STATUS = ?, PRIORITY = ? where id = ?",
		task.Title, task.Desc, task.Status, task.Priority, task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) Delete(c context.Context, s string) error {
	_, err := r.db.ExecContext(c, "DELETE FROM tasks WHERE ID = ?", s)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) Get(c context.Context) ([]entity.Task, error) {
	rows, err := r.db.QueryContext(c, "select * from tasks")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	tasks := []entity.Task{}

	for rows.Next() {
		t := entity.Task{}
		err := rows.Scan(&t.ID, &t.Title, &t.Desc, &t.Status, &t.Priority)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tasks = append(tasks, t)
	}
	return tasks, err
}

func UpMigrations(db *sql.DB) error {
	if err := goose.SetDialect("sqlite"); err != nil {
		return err
	}
	if err := goose.Up(db, "."); err != nil {
		return err
	}
	return nil
}
