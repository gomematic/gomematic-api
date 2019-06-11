package model

import (
	"time"
)

// Team within Gomematic.
type Team struct {
	ID        string `storm:"id"`
	Slug      string `storm:"index,unique"`
	Name      string `storm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
