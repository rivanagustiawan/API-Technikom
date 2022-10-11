package user

type Service interface {
	GetUser() ([]User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUser() ([]User, error) {
	users, err := s.repository.FindAll()

	if err != nil {
		return users, err
	}
	return users, nil

}
