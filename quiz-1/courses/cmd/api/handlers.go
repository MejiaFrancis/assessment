// Filename: cmd/api/courses.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MejiaFrancis/assesment/quiz-1/courses/internal/data"
)

func (app *application) createCourseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created a course...")
}

func (app *application) showCourseHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	//fmt.Fprintf(w, "show details of courses.go %d\n", id)

	course := data.Course{
		Course_ID:            id,
		Course_Code:          "code",
		Course_NumberCredits: 105,
		CreateAt:             time.Now(),
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
