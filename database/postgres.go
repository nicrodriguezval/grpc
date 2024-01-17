package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nicrodriguezval/grpc/models"
	"log"
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

func (r *PostgresRepository) Enroll(ctx context.Context, enrollment *models.Enrollment) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO enrollments (id, student_id, test_id) VALUES ($1, $2, $3)", enrollment.Id, enrollment.StudentId, enrollment.TestId)
	return err
}

func (r *PostgresRepository) GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT s.id, s.name, s.age FROM students s INNER JOIN enrollments e ON s.id = e.student_id WHERE e.test_id = $1", testId)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	students := make([]*models.Student, 0)
	for rows.Next() {
		student := models.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	return students, nil
}

func (r *PostgresRepository) GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, question FROM questions WHERE test_id = $1", testId)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	questions := make([]*models.Question, 0)
	for rows.Next() {
		question := models.Question{}
		err := rows.Scan(&question.Id, &question.Question)
		if err != nil {
			return nil, err
		}
		questions = append(questions, &question)
	}

	return questions, nil
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}
