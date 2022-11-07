package invasion

import (
	"fmt"

	petname "github.com/dustinkirkland/golang-petname"
)

type Alien struct {
	Id   int
	City *City
	name string
}

func NewAlien(id int, friendly bool) *Alien {
	var name string
	if friendly {
		name = petname.Generate(3, "-")
	} else {
		name = fmt.Sprintf("alien %d", id+1)
	}

	return &Alien{id, nil, name}
}

func (a *Alien) Name() string {
	return a.name
}
