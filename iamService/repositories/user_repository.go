package repositories

import (
	"IAMService/models"

	"github.com/google/uuid"
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

func (r *UserRepository) Create(user *models.User) error {
	var err error
	user.ID, err = uuid.NewRandom()
	if err != nil {
		return err
	}
	err = r.db.Create(user).Error
	return err
}

func (r *UserRepository) FindByID(id uuid.UUID) *models.User {
	var user models.User
	r.db.First(&user, id)

	if user == (models.User{}) {
		return nil
	}

	return &user
}

func (r *UserRepository) FindByEmail(email string) *models.User {
	var user models.User
	r.db.Where("email = ?", email).First(&user)

	if user == (models.User{}) {
		return nil
	}

	return &user
}
