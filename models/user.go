package models

import "time"

type User struct {
	id      int
	name    string
	created time.Time
	status  bool
}

func (this *User) AddUser(id int, name string, created time.Time, status bool) {
	this.id = id
	this.name = name
	this.created = created
	this.status = status
}
