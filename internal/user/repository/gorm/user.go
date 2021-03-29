package gorm

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/CaninoDev/gastro/server/api/user"
	"github.com/CaninoDev/gastro/server/internal/logger"
)

// userRepository provides persistent repository for user
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository instantiates an instance for data persistence
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) View(ctx context.Context, user *user.User) error {
	if err := r.db.First(&user, user.ID).Error; errors.Is(
		err,
		gorm.ErrRecordNotFound) {
		return err
	} else if err != nil {
		logger.Error.Printf("db connection error %v", err)
		return err
	}
	return nil
}

func (r *userRepository) Search(ctx context.Context, user *user.User) error {
	if err := r.db.Where("email = ? ", user.Email).Or("first_name = ? AND last_name = ?",
		user.FirstName, user.LastName).Or("id = ?", user.ID).First(&user).Error; errors.Is(
		err,
		gorm.ErrRecordNotFound) {
		return err
	} else if err != nil {
		logger.Error.Printf("db connection error %v", err)
		return err
	}
	return nil
}

func (r *userRepository) Create(ctx context.Context, user *user.User) error {
	return r.db.Create(&user).Error
}
func (r *userRepository) Update(ctx context.Context, user *user.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) Delete(ctx context.Context, user *user.User) error {
	return r.db.Delete(&user, "id = ?", user.ID).Error
}