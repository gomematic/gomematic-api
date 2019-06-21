package model

import (
	"time"
)

// Team within Gomematic.
type Team struct {
	ID        string `storm:"id" gorm:"primary_key"`
	Slug      string `storm:"unique" sql:"unique_index"`
	Name      string `storm:"unique" sql:"unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []*TeamUser
}
