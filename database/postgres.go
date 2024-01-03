package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nicrodriguezval/grpc/models"
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

func (r *PostgresRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)
	return err
}

func (r *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	student := models.Student{}
	err := row.Scan(&student.Id, &student.Name, &student.Age)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *PostgresRepository) CreateTest(ctx context.Context, test *models.Test) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO tests (id, name) VALUES ($1, $2)", test.Id, test.Name)
	return err
}

func (r *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id)
	test := models.Test{}
	err := row.Scan(&test.Id, &test.Name)
	if err != nil {
		return nil, err
	}

	return &test, nil
}

func (r *PostgresRepository) CreateQuestion(ctx context.Context, question *models.Question) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO questions (id, test_id, question, answer) VALUES ($1, $2, $3, $4)", question.Id, question.TestId, question.Question, question.Answer)
	return err
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}
