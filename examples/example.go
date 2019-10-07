package examples

import "time"

// gen:ddd {"table": "examples"}
type Example struct {
	Id           *int       `db:"field=id,primary"`
	Name         string     `db:"field=name"`
	Phone        string     `db:"field=phone,searchable"`
	Email        string     `db:"field=email"`
	PasswordHash string     `db:"field=passwordHash"`
	Created      time.Time  `db:"field=created,searchGroup=Timestamps"`
	Updated      time.Time  `db:"field=updated,searchGroup=Timestamps"`
	Activated    *time.Time `db:"field=activated"`
	Deleted      *time.Time `db:"field=deleted,active"`
}
