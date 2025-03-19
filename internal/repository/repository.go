package repository

import (
	"database/sql"
	"fmt"
	"go/rest/internal/app/files"
	"go/rest/internal/entity"

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

	initQuery := `CREATE TABLE IF NOT EXISTS tasks (
id TEXT PRIMARY KEY,
title TEXT NOT NULL,
description TEXT,
status BOOLEAN,
priority TEXT);`

	_, err = db.Exec(initQuery)
	if err != nil {
		panic(err)
	}
	return &Repo{db: db}
}

func (r *Repo) Create(task entity.Task) error {
	query := `INSERT INTO tasks (ID, TITLE, DESCRIPTION, STATUS, PRIORITY) VALUES (?, ?, ?, ?, ?);`

	_, err := r.db.Exec(query, task.ID, task.Title, task.Desc, task.Status, task.Priority)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) Update(task entity.Task) error {
	_, err := r.db.Exec("UPDATE tasks SET TITLE=?, DESCRIPTION=?, STATUS = ?, PRIORITY = ? where id = ?",
		task.Title, task.Desc, task.Status, task.Priority, task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) Delete(s string) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE ID = ?", s)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) Get() ([]entity.Task, error) {
	rows, err := r.db.Query("select * from tasks")
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
