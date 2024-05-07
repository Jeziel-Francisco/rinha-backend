package entities

import "reflect"

type GetPersonByTermIntention struct {
	Term   string
	People []*Person
}

func (intention *GetPersonByTermIntention) IsEmpty() bool {
	return reflect.DeepEqual(&GetPersonByTermIntention{}, intention)
}
