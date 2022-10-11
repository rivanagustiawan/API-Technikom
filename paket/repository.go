package paket

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Paket, error)
	FindById(ID int) (Paket, error)
	Save(paket Paket) (Paket, error)
	SaveDesk(deskPaket Deskripsi_Paket) (Deskripsi_Paket, error)
	Update(paket Paket) (Paket, error)
	Delete(ID int)
	DeleteDesk(ID int)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Paket, error) {
	var pakets []Paket

	err := r.db.Preload("Deskripsi").Find(&pakets).Error

	if err != nil {
		return pakets, err
	}

	return pakets, nil
}
func (r *repository) FindById(ID int) (Paket, error) {
	var paket Paket

	err := r.db.Where("ID", ID).Preload("Deskripsi").Find(&paket).Error

	if err != nil {
		return paket, err
	}

	return paket, nil
}

func (r *repository) Save(paket Paket) (Paket, error) {
	err := r.db.Create(&paket).Error
	if err != nil {
		return paket, err
	}
	return paket, nil
}

func (r *repository) SaveDesk(desk Deskripsi_Paket) (Deskripsi_Paket, error) {

	err := r.db.Create(&desk).Error
	if err != nil {
		return desk, err
	}
	return desk, nil
}

func (r *repository) Update(paket Paket) (Paket, error) {
	err := r.db.Save(&paket).Error
	if err != nil {
		return paket, err
	}
	return paket, nil
}
func (r *repository) Delete(ID int) {
	var paket Paket
	var desk Deskripsi_Paket
	r.db.Delete(&paket, ID)
	r.db.Where("Id_Paket", ID).Delete(&desk)
}
func (r *repository) DeleteDesk(ID int) {
	var desk Deskripsi_Paket
	r.db.Delete(&desk, ID)
}
