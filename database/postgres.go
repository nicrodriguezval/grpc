package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"grpc/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (r *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO students (name, age) VALUES ($1, $2)", student.Name, student.Age)
	return err
}

func (r *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	student := &models.Student{}
	err := row.Scan(student.Name, student.Age)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}
