// Filename: ./internal/data/courses.go
package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq"
)

// course represents one row of data in our schools table
type Course struct { //we can get data from client and put it in here and send to db or vise versa
	ID           int64     `json:"id"`
	CourseCode   string    `json:"course_code"`
	CourseTitle  string    `json:"course_title"`
	CourseCredit string    `json:"course_credit"`
	CreateAt     time.Time `json:"-"`
	Version      int32     `json:"version"`
}

type CourseModel struct {
	DB *sql.DB
}

// Insert() allows us  to create a new Course
func (m CourseModel) Insert(course *Course) error {
	query := `
		INSERT INTO courses (course_code, course_title, course_credit)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, version
	`
	// Collect the data fields into a slice
	args := []interface{}{
		course.CourseCode,
		course.CourseTitle,
		course.CourseCredit,
	}
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&course.ID, &course.CreateAt, &course.Version)
}

// Get() allows us to retrieve a specific School
func (m CourseModel) Get(id int64) (*Course, error) {
	// Ensure that there is a valid id
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Create the query
	query := `
		SELECT id, created_at, coursecode, coursetitle, coursecredit, version
		FROM courses
		WHERE id = $1
	`
	// Declare a School variable to hold the returned data
	var course Course
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Execute the query using QueryRow()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&course.ID,
		&course.CreateAt,
		&course.CourseCode,
		&course.CourseTitle,
		&course.CourseCredit,
		&course.Version,
	)
	// Handle any errors
	if err != nil {
		// Check the type of error
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	// Success
	return &course, nil
}

// Update() allows us to edit/alter a specific Course
// Optimistic locking (version number)
func (m CourseModel) Update(course *Course) error {
	// Create a query
	query := `
		UPDATE courses
		SET cousecode = $1, coursetitle = $2, coursecredit= $3,
		     version = version + 1
		WHERE id = $4
		AND version = $5
		RETURNING version
	`
	args := []interface{}{
		course.CourseCode,
		course.CourseTitle,
		course.CourseCredit,
		course.ID,
		course.Version,
	}
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Check for edit conflicts
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&course.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

// Delete() removes a specific Course
func (m CourseModel) Delete(id int64) error {
	// Ensure that there is a valid id
	if id < 1 {
		return ErrRecordNotFound
	}
	// Create the delete query
	query := `
		DELETE FROM courses
		WHERE id = $1
	`
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Execute the query
	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	// Check how many rows were affected by the delete operation. We
	// call the RowsAffected() method on the result variable
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// Check if no rows were affected
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
