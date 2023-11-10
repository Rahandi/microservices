package repositories

import (
	"iamService/models"

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

func (r *UserRepository) Create(user *models.DBUser) error {
	var err error
	user.ID, err = uuid.NewRandom()
	if err != nil {
		return err
	}
	err = r.db.Create(user).Error
	return err
}

func (r *UserRepository) FindByID(id uuid.UUID) *models.DBUser {
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
