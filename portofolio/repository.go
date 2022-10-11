package portofolio

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Portofolio, error)
	FindById(ID int) (Portofolio, error)
	Save(porto Portofolio) (Portofolio, error)
	Update(Portofolio Portofolio) (Portofolio, error)
	Delete(ID int)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Portofolio, error) {
	var porto []Portofolio

	err := r.db.Preload("Klien").Find(&porto).Error

	if err != nil {
		return porto, err
	}

	return porto, nil
}
func (r *repository) FindById(ID int) (Portofolio, error) {
	var porto Portofolio

	err := r.db.Where("ID", ID).Find(&porto).Error

	if err != nil {
		return porto, err
	}

	return porto, nil
}

func (r *repository) Save(porto Portofolio) (Portofolio, error) {
	err := r.db.Create(&porto).Error
	if err != nil {
		return porto, err
	}
	return porto, nil
}

func (r *repository) Update(porto Portofolio) (Portofolio, error) {
	err := r.db.Save(&porto).Error
	if err != nil {
		return porto, err
	}
	return porto, nil
}
func (r *repository) Delete(ID int) {
	var porto Portofolio

	r.db.Delete(&porto, ID)
}
