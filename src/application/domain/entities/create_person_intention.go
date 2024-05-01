package entities

import "reflect"

type CreatePersonIntention struct {
	Person Person
}

func (intention *CreatePersonIntention) IsEmpty() bool {
	return reflect.DeepEqual(&CreatePersonIntention{}, intention)
}
