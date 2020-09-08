package models

// Character struct model
type Character struct {
	ID     int64  `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Gender string `db:"gender" json:"gender"`
	Height string `db:"height" json:"height"`
	Class  string `db:"class" json:"class"`
}
