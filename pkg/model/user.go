package model

import (
	"time"
)

// User within Gomematic.
type User struct {
	ID        string `storm:"id"`
	Slug      string `storm:"index,unique"`
	Username  string `storm:"unique"`
	Password  string
	Email     string `storm:"unique"`
	Active    bool
	Admin     bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
