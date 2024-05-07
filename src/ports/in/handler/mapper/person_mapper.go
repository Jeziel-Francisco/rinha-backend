package mapper

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/ports/in/handler/dto"
)

func FromRequestCreatePersonDtoToCreatePersonIntention(source *dto.RequestCreatePersonDto) *entities.CreatePersonIntention {
	if source == nil || source.IsEmpty() {
		return nil
	}
	return &entities.CreatePersonIntention{
		Person: entities.Person{
			Nickname:  source.Nickname,
			Name:      source.Name,
			BirthDate: source.BirthDate,
			Stacks:    source.Stacks,
		},
	}
}

func FromPersonIDToResponseCreatePersonDto(personID *string) *dto.ResponseCreatePersonDto {
	return &dto.ResponseCreatePersonDto{
		ID: *personID,
	}
}

func FromRequestGetPersonByIDToCreatePersonIntention(source *dto.RequestGetPersonByID) *entities.GetPersonByIDIntention {
	if source == nil || source.IsEmpty() {
		return nil
	}
	return &entities.GetPersonByIDIntention{
		Person: entities.Person{
			ID: source.ID,
		},
	}
}

func FromGetPersonByIDToResponseGetPersonDetail(source *entities.Person) *dto.ResponseGetPersonDetail {
	if source == nil || source.IsEmpty() {
		return nil
	}
	return &dto.ResponseGetPersonDetail{
		ID:        source.ID,
		Nickname:  source.Nickname,
		Name:      source.Name,
		BirthDate: source.BirthDate,
		Stacks:    source.Stacks,
	}
}

func FromRequestGetPersonByTermToCreatePersonIntention(source *dto.RequestGetPersonByTerm) *entities.GetPersonByTermIntention {
	if source == nil || source.IsEmpty() {
		return nil
	}
	return &entities.GetPersonByTermIntention{
		Term: source.Term,
	}
}

func FromGetPersonByTermToListResponseGetPersonDetail(sources []*entities.Person) []*dto.ResponseGetPersonDetail {
	if len(sources) == 0 {
		return []*dto.ResponseGetPersonDetail{}
	}

	var result []*dto.ResponseGetPersonDetail = []*dto.ResponseGetPersonDetail{}

	for _, source := range sources {
		if source.IsEmpty() {
			continue
		}
		result = append(result, FromGetPersonByIDToResponseGetPersonDetail(source))
	}
	return result
}
