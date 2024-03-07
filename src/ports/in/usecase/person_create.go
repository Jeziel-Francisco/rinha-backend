package usecase

type PersonCreate interface {
	Execute() (err error)
}
