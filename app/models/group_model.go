package models

import "github.com/google/uuid"

type Group struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Title        string    `db:"title" json:"title"`
	Descriptions string    `db:"descriptions" json:"descriptions"`
}
