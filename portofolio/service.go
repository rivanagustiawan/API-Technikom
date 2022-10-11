package portofolio

type Service interface {
	GetPortofolios() ([]Portofolio, error)
	GetPortofolio(input GetIdPorto) (Portofolio, error)
	CreatePortofolio(input PortofolioInput) (Portofolio, error)
	UpdatePortofolio(inputID GetIdPorto, input PortofolioInput) (Portofolio, error)
	DeletePortofolio(input GetIdPorto)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetPortofolios() ([]Portofolio, error) {
	porto, err := s.repository.FindAll()

	if err != nil {
		return porto, err
	}
	return porto, nil

}
func (s *service) GetPortofolio(input GetIdPorto) (Portofolio, error) {

	porto, err := s.repository.FindById(input.ID)

	if err != nil {
		return porto, err
	}
	return porto, nil

}

func (s *service) CreatePortofolio(input PortofolioInput) (Portofolio, error) {
	porto := Portofolio{}
	porto.NamaApps = input.NamaApps
	porto.Id_Klien = input.Id_Klien
	porto.JenisPaket = input.JenisPaket
	porto.LinkApps = input.LinkApps
	porto.DeskripsiApps = input.DeskripsiApps
	porto.PhotoApps1 = input.PhotoApps1
	porto.PhotoApps2 = input.PhotoApps2
	porto.PhotoApps3 = input.PhotoApps3
	porto.PhotoApps4 = input.PhotoApps4

	newPorto, err := s.repository.Save(porto)
	if err != nil {
		return newPorto, err
	}
	return newPorto, nil
}

func (s *service) UpdatePortofolio(inputID GetIdPorto, input PortofolioInput) (Portofolio, error) {

	porto, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return porto, err
	}
	porto.NamaApps = input.NamaApps
	porto.Id_Klien = input.Id_Klien
	porto.JenisPaket = input.JenisPaket
	porto.LinkApps = input.LinkApps
	porto.DeskripsiApps = input.DeskripsiApps
	porto.PhotoApps1 = input.PhotoApps1
	porto.PhotoApps2 = input.PhotoApps2
	porto.PhotoApps3 = input.PhotoApps3
	porto.PhotoApps4 = input.PhotoApps4

	updatePorto, err := s.repository.Update(porto)
	if err != nil {
		return updatePorto, err
	}
	return updatePorto, nil
}

func (s *service) DeletePortofolio(input GetIdPorto) {

	s.repository.Delete(input.ID)

}
