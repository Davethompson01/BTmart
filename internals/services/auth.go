package services

import (
	"bsmart/config"
	"bsmart/internals/models"
	"bsmart/internals/repositary"
	Validation "bsmart/internals/validation"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(config config.Apiconfig, auth models.MartAuth) (int, error) {

	// validate Input data
	err := Validation.CreateAccount(auth)
	if err != nil {
		return 0, fmt.Errorf("validate account input; %w", err)
	}

	// check if mail exists
	checkMailExist := repositary.CheckMailExist(config, auth.Email)
	if !checkMailExist {
		return 0, errors.New("email exist already")
	}

	// create hashpassword
	hash, err := HashPassword(auth.Password)
	if err != nil {
		return 0, fmt.Errorf("hash password failed %v", err)
	}

	auth.Password = hash
	auth.Auth_method = "Provider"

	createAccount, err := repositary.CreateAccount(config, auth)
	if err != nil {
		return 0, fmt.Errorf("Failed to create account: %w", createAccount)
	}

	config.Logger.Info(
		"account created successfully",
		"user_id", createAccount,
		"auth_method", auth.Auth_method,
	)

	return createAccount, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// ComparePassword compares a hashed password with a plain-text password.
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}