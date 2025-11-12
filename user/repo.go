package user

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type Repo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	);`
	_, err := r.DB.Exec(query)
	return err
}

func (r *Repo) InsertUser(name string) (int, error) {
	var id int
	err := r.DB.QueryRow(`INSERT INTO users (name) VALUES ($1) RETURNING id`, name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insert failed: %w", err)
	}
	return id, nil
}

func (r *Repo) GetUser(id int) (User, error) {
	var u User
	err := r.DB.QueryRow(`SELECT id, name FROM users WHERE id=$1`, id).Scan(&u.ID, &u.Name)
	if err != nil {
		return User{}, fmt.Errorf("get failed: %w", err)
	}
	return u, nil
}
