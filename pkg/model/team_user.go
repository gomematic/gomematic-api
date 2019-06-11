package model

import (
	"time"
)

// TeamUser within Gomematic.
type TeamUser struct {
	TeamID    string `storm:"id,index"`
	Team      *Team
	UserID    string `storm:"id,index"`
	User      *User
	Perm      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
