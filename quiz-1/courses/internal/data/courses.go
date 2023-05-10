package data

import (
	"time"
)

// School represents one row of data in our schools table
type Course struct {
	Course_ID            int64     `json:"id"`
	Course_Code          string    `json:"code"`
	Course_NumberCredits int       `json:"credits"`
	CreateAt             time.Time `json:"-"`
}
