package repository

import (
	"bookstore/internal/models"
	"database/sql"
	"errors"
)

// 2. UserRepository Struct
type UserRepository struct {
	DB *sql.DB
}

// 3. Constructor Function: NewUserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// 4. CreateUser Method
func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, created_at"
	err := r.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
	return err
}

// 5. GetUserByEmail Method
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, name, email, password, created_at FROM users WHERE email=$1"
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	return user, err
}

// Новый метод: Получить пользователя по ID
func (r *UserRepository) GetUserByID(userID int) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, name, email FROM users WHERE id = $1"
	err := r.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
