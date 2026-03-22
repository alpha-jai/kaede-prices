package repo

import (
	"github.com/koyo/kaede-prices/domain"
)

type userRepository struct {
	// db *sql.DB
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *domain.User) error {
	// Implementation here
	return nil
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	// Implementation here
	return &domain.User{ID: 1, Email: email, Password: "password"}, nil
}

func (r *userRepository) GetByID(id uint) (*domain.User, error) {
	// Implementation here
	return &domain.User{ID: id, Email: "test@example.com"}, nil
}
