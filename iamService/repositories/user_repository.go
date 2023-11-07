package repositories

import (
	"iamService/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	userRepository := &UserRepository{
		db: db,
	}
	return userRepository
}

func (r *UserRepository) Create(user *models.DBUser) error {
	err := r.db.Create(user).Error
	return err
}

func (r *UserRepository) FindByID(id uint) *models.DBUser {
	var user models.DBUser
	r.db.First(&user, id)

	if user == (models.DBUser{}) {
		return nil
	}

	return &user
}

func (r *UserRepository) FindByEmail(email string) *models.DBUser {
	var user models.DBUser
	r.db.Where("email = ?", email).First(&user)

	if user == (models.DBUser{}) {
		return nil
	}

	return &user
}
