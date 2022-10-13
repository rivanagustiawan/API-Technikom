package klien

type Service interface {
	GetKliens() ([]Klien, error)
	GetKlien(input GetIdKlien) (Klien, error)
	CreateKlien(input KlienInput, path string) (Klien, error)
	UpdateKlien(inputID GetIdKlien, input KlienInputUpdate, path string) (Klien, error)
	DeleteKlien(input GetIdKlien)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetKliens() ([]Klien, error) {
	kliens, err := s.repository.FindAll()

	if err != nil {
		return kliens, err
	}
	return kliens, nil

}
func (s *service) GetKlien(input GetIdKlien) (Klien, error) {

	klien, err := s.repository.FindById(input.ID)

	if err != nil {
		return klien, err
	}
	return klien, nil

}

func (s *service) CreateKlien(input KlienInput, path string) (Klien, error) {
	klien := Klien{}
	klien.NamaKlien = input.NamaKlien
	klien.PicLogo = path

	newKlien, err := s.repository.Save(klien)
	if err != nil {
		return newKlien, err
	}
	return newKlien, nil
}

func (s *service) UpdateKlien(inputID GetIdKlien, input KlienInputUpdate, path string) (Klien, error) {

	klien, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return klien, err
	}
	klien.NamaKlien = input.NamaKlien
	// klien.PicLogo = input.PicLogo

	updateKlien, err := s.repository.Update(klien)
	if err != nil {
		return updateKlien, err
	}
	return updateKlien, nil
}

func (s *service) DeleteKlien(input GetIdKlien) {

	s.repository.Delete(input.ID)

}
