package user

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]User, error)
	// FindById(ID int) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}
