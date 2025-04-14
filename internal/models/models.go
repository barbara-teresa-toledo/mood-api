package models

import "time"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type MoodEntry struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Level     int       `json:"level"`
	Timestamp time.Time `json:"timestamp"`
}

type MoodReport struct {
	UserID     string
	Average    float64
	From       time.Time
	To         time.Time
	EntryCount int
}
