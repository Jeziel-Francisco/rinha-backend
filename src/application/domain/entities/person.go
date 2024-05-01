package entities

import "reflect"

type Person struct {
	ID        string
	Nickname  string
	Name      string
	BirthDate string
	Stacks    []string
}

func (entity *Person) IsEmpty() bool {
	return reflect.DeepEqual(&Person{}, entity)
}
