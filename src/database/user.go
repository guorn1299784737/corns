package database

import (
	"github.com/asdine/storm/v3"
	"time"
)

type User struct {
	ID int // primary key
	Group string `storm:"index"` // this field will be indexed
	Email string `storm:"unique"` // this field will be indexed with a unique constraint
	Name string // this field will not be indexed
	Age int `storm:"index"`
	CreatedAt time.Time
}

var model = User{
	ID: 10,
	Group: "staff",
	Email: "john@provider.com",
	Name: "John",
	Age: 21,
	CreatedAt: time.Now(),
}

func init() {
	db, err := storm.Open("my.db")
	defer db.Close()
}

