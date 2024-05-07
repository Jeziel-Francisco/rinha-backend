package entities

import "reflect"

type GetPersonByIDIntention struct {
	Person Person
}

func (intention *GetPersonByIDIntention) IsEmpty() bool {
	return reflect.DeepEqual(&GetPersonByIDIntention{}, intention)
}
