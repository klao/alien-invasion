package invasion

import "fmt"

type Alien struct {
	Id   int
	City *City
}

func (a *Alien) Name() string {
	return fmt.Sprintf("alien %d", a.Id+1)
}
