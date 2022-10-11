package paket

type Service interface {
	GetPakets() ([]Paket, error)
	GetPaket(input GetIdPaket) (Paket, error)
	CreatePaket(input PaketInput) (Paket, error)
	CreateDeskPaket(input Deskripsi_PaketInput) (Paket, error)
	Update(inputID GetIdPaket, input PaketInput) (Paket, error)
	DeletePaket(input GetIdPaket)
	DeleteDeskripsi(input GetIdPaket)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetPakets() ([]Paket, error) {
	pakets, err := s.repository.FindAll()

	if err != nil {
		return pakets, err
	}
	return pakets, nil

}
func (s *service) GetPaket(input GetIdPaket) (Paket, error) {

	paket, err := s.repository.FindById(input.ID)

	if err != nil {
		return paket, err
	}
	return paket, nil

}

func (s *service) CreatePaket(input PaketInput) (Paket, error) {
	paket := Paket{}
	paket.NamaPaket = input.NamaPaket
	paket.Harga = input.Harga
	paket.JenisPaket = input.JenisPaket
	paket.WarnaLabel = input.WarnaLabel

	newPaket, err := s.repository.Save(paket)
	if err != nil {
		return newPaket, err
	}
	return newPaket, nil
}
func (s *service) CreateDeskPaket(input Deskripsi_PaketInput) (Paket, error) {
	desk := Deskripsi_Paket{}
	desk.Id_Paket = input.Id_Paket
	desk.Deskripsi = input.Deskripsi

	s.repository.SaveDesk(desk)

	paket, _ := s.repository.FindById(input.Id_Paket)
	return paket, nil
}

func (s *service) Update(inputID GetIdPaket, input PaketInput) (Paket, error) {

	paket, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return paket, err
	}
	paket.NamaPaket = input.NamaPaket
	paket.Harga = input.Harga
	paket.JenisPaket = input.JenisPaket
	paket.WarnaLabel = input.WarnaLabel

	updatePaket, err := s.repository.Update(paket)
	if err != nil {
		return updatePaket, err
	}
	return updatePaket, nil
}

func (s *service) DeletePaket(input GetIdPaket) {

	s.repository.Delete(input.ID)

}
func (s *service) DeleteDeskripsi(input GetIdPaket) {

	s.repository.DeleteDesk(input.ID)

}
