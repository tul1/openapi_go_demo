package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
	"github.com/tul1/openapi_go_demo/model"
)

type SQLHandler struct {
	db *sql.DB
}

// NewSQLHandler initializes a new SQLHandler with a SQLite database.
func NewSQLHandler(dsn string) (*SQLHandler, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	// Create the todos table if it doesn't exist.
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL
	)`)
	if err != nil {
		return nil, err
	}

	return &SQLHandler{db: db}, nil
}

func (s *SQLHandler) GetTodos() ([]model.Todo, error) {
	rows, err := s.db.Query("SELECT id, task FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Task); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (s *SQLHandler) AddTodo(todo model.Todo) error {
	_, err := s.db.Exec("INSERT INTO todos (task) VALUES (?)", todo.Task)
	return err
}

func (s *SQLHandler) GetTodoByID(id int) (model.Todo, error) {
	var todo model.Todo
	err := s.db.QueryRow("SELECT id, task FROM todos WHERE id = ?", id).Scan(&todo.ID, &todo.Task)
	if err == sql.ErrNoRows {
		return todo, fmt.Errorf("todo not found")
	}
	return todo, err
}

func (s *SQLHandler) DeleteTodoByID(id int) error {
	_, err := s.db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
