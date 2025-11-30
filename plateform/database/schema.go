package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string
	Email     string
	Validated bool
}

type ActivationCode struct {
	gorm.Model
	Code       string
	Expiration time.Time
	UserID     int
	User       User
}

type Contest struct {
	gorm.Model
	Name      string
	StartTime time.Time
	EndTime   time.Time
}

type Difficulty struct {
	gorm.Model
	Name string
}

type Problem struct {
	gorm.Model
	Name         string
	Points       int
	DifficultyID int
	Difficulty   Difficulty
}
