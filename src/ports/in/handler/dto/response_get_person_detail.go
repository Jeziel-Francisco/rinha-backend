package dto

type ResponseGetPersonDetail struct {
	ID        string   `json:"id"`
	Nickname  string   `json:"apelido"`
	Name      string   `json:"nome"`
	BirthDate string   `json:"nascimento"`
	Stacks    []string `json:"stack"`
}
