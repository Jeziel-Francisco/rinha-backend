package entities

import "reflect"

type CountPersonIntention struct {
	QuantityPerson uint64
}

func (intention *CountPersonIntention) IsEmpty() bool {
	return reflect.DeepEqual(&CountPersonIntention{}, intention)
}
