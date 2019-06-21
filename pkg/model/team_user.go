package model

import (
	"time"
)

// TeamUser within Gomematic.
type TeamUser struct {
	TeamID    string `storm:"id,index" sql:"index"`
	Team      *Team
	UserID    string `storm:"id,index" sql:"index"`
	User      *User
	Perm      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
