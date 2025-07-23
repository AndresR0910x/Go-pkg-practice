package models

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	Email     string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (this User) AddUser(id int, name string, email string, status bool) {
	this.Id = id
	this.Name = name
	this.Email = email
	this.Status = status
	this.CreatedAt = time.Now()
	this.UpdatedAt = time.Now()
}
