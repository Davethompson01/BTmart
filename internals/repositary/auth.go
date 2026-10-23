package repositary

import (
	"bsmart/config"
	"bsmart/internals/models"
	"context"
	"time"
)

func CreateAccount(config config.Apiconfig, auth models.MartAuth) (int, error) {
	query := `INSERT INTO users(fullname, email, number, school, password, profile_pic) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING user_id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var userID int
	err := config.DB.QueryRowContext(ctx, query, auth.Fullname, auth.Email, auth.PhoneNumber, auth.Auth_method, auth.School, auth.Password, auth.Picture).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func CheckMailExist(config config.Apiconfig, email string) bool {
	query := `SELECT EXISTS (
		SELECT 1 FROM users WHERE email = $1
	)`
	var exists bool

	err := config.DB.QueryRow(query, exists).Scan(&email)
	if err != nil {
		return false
	}
	return true
}
