package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"go/rest/internal/app/files"
	"go/rest/internal/entity"
	_ "go/rest/migrations"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

type Repo struct {
	db *sql.DB
}

func New() *Repo {
	data := files.OpenYaml()
	host := data["host"]
	port := data["postgresPort"]
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic("Can`t parse port")
	}
	user := data["postgresUser"]
	password := data["postgresPass"]
	dbname := data["postgresDB"]

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, portInt, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	query := `CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		status BOOLEAN,
		priority TEXT);`

	_, err = db.Exec(query)
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
	query := `INSERT INTO tasks (ID, TITLE, DESCRIPTION, STATUS, PRIORITY) VALUES ($1, $2, $3, $4, $5);`

	_, err := r.db.ExecContext(c, query, task.ID, task.Title, task.Desc, task.Status, task.Priority)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) Update(c context.Context, task entity.Task) error {
	_, err := r.db.ExecContext(c, "UPDATE tasks SET TITLE=$1, DESCRIPTION=$2, STATUS = $3, PRIORITY = $4 where id = $5",
		task.Title, task.Desc, task.Status, task.Priority, task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) Delete(c context.Context, s string) error {
	_, err := r.db.ExecContext(c, "DELETE FROM tasks WHERE ID = $1", s)
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
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(db, "."); err != nil {
		return err
	}
	return nil
}
