package farm

import (
	"fmt"
)

type Farmer struct {
	ID   int
	Name string
}

func (f Farmer) FeedAnimal(a Animal) string {
	return fmt.Sprintf("%s is feeding the %s named %s.", f.Name, a.GetType(), a.GetName())
}
