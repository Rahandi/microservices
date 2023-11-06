package repositories

import (
	"iamService/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	authRepository := &AuthRepository{
		db: db,
	}
	return authRepository
}

func (r *AuthRepository) Create(user *models.User) error {
	err := r.db.Create(user).Error
	return err
}

func (r *AuthRepository) FindByID(id uint) *models.User {
	var user models.User
	r.db.First(&user, id)

	if user == (models.User{}) {
		return nil
	}

	return &user
}

func (r *AuthRepository) FindByEmail(email string) *models.User {
	var user models.User
	r.db.Where("email = ?", email).First(&user)

	if user == (models.User{}) {
		return nil
	}

	return &user
}
