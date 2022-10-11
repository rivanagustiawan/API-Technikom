package klien

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Klien, error)
	FindById(ID int) (Klien, error)
	Save(klien Klien) (Klien, error)
	Update(klien Klien) (Klien, error)
	Delete(ID int)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Klien, error) {
	var kliens []Klien

	err := r.db.Find(&kliens).Error

	if err != nil {
		return kliens, err
	}

	return kliens, nil
}
func (r *repository) FindById(ID int) (Klien, error) {
	var klien Klien

	err := r.db.Where("ID", ID).Find(&klien).Error

	if err != nil {
		return klien, err
	}

	return klien, nil
}

func (r *repository) Save(klien Klien) (Klien, error) {
	err := r.db.Create(&klien).Error
	if err != nil {
		return klien, err
	}
	return klien, nil
}

func (r *repository) Update(klien Klien) (Klien, error) {
	err := r.db.Save(&klien).Error
	if err != nil {
		return klien, err
	}
	return klien, nil
}
func (r *repository) Delete(ID int) {
	var klien Klien

	r.db.Delete(&klien, ID)
}
