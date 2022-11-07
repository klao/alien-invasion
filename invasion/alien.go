package invasion

import "fmt"

type Alien struct {
	Id   int
	City *City
}

func (a *Alien) Name() string {
	return fmt.Sprint(a.Id + 1)
}
