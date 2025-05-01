package repository

import (
	"context"
	"fmt"
	"go_clean_architecture/Infrastructure/middleware"
	"go_clean_architecture/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// signup repository

type signUpRepository struct {
	database        mongo.Database
	collection      string
	passwordService *middleware.PasswordService
}

func NewSignUpRepository(database mongo.Database, collection string, passwordService *middleware.PasswordService) domain.SignUpRepository {
	return &signUpRepository{
		database:        database,
		collection:      collection,
		passwordService: passwordService,
	}
}

// SignUp implements domain.SignUpRepository.
func (s *signUpRepository) SignUp(c context.Context, signUpRequest domain.SignUpRequest) (domain.SignUpResponse, error) {
	// Check if user already exists
	collection := s.database.Collection(s.collection)
	var existingUser domain.SignUpRequest
	err := collection.FindOne(c, bson.M{"username": signUpRequest.Username}).Decode(&existingUser)
	if err == nil {
		return domain.SignUpResponse{}, fmt.Errorf("user with username %s already exists", signUpRequest.Username)
	}
	if err != mongo.ErrNoDocuments {
		return domain.SignUpResponse{}, err
	}

	hashedPassword, err := s.passwordService.HashPassword(signUpRequest.Password)
	if err != nil {
		return domain.SignUpResponse{}, err
	}

	signUpRequest.Password = hashedPassword

	_, err = collection.InsertOne(c, signUpRequest)
	if err != nil {
		return domain.SignUpResponse{}, err
	}
	return domain.SignUpResponse{
		Message: "Signup successful",
	}, nil
}

// login repository

type loginRepository struct {
	database        mongo.Database
	collection      string
	passwordService *middleware.PasswordService
	jwtService      *middleware.JWTService
}

func NewLoginRepository(database mongo.Database, collection string, passwordService *middleware.PasswordService, jwtService *middleware.JWTService) domain.LoginRepository {
	return &loginRepository{
		database:        database,
		collection:      collection,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

// Login implements domain.LoginRepository.
func (l *loginRepository) Login(c context.Context, loginRequest domain.LoginRequest) (domain.LoginResponse, error) {
	collection := l.database.Collection(l.collection)

	var signUpRequest domain.SignUpRequest
	err := collection.FindOne(c, bson.M{"username": loginRequest.Username}).Decode(&signUpRequest)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	if err := l.passwordService.VerifyPassword(signUpRequest.Password, loginRequest.Password); err != nil {
		return domain.LoginResponse{}, err
	}

	tokenString, err := l.jwtService.GenerateToken(signUpRequest.Username, signUpRequest.Role)
	if err != nil {
		fmt.Printf("Error generating token: %v\n", err)
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		Message: "Login successful as " + signUpRequest.Role,
		Token:   tokenString,
	}, nil
}

// refresh token repository

type refreshTokenRepository struct {
	database   mongo.Database
	collection string
	jwtService *middleware.JWTService
}

func NewRefreshTokenRepository(database mongo.Database, collection string, jwtService *middleware.JWTService) domain.RefreshTokenRepository {
	return &refreshTokenRepository{
		database:   database,
		collection: collection,
		jwtService: jwtService,
	}
}

// RefreshToken implements domain.RefreshTokenRepository.
func (r *refreshTokenRepository) RefreshToken(c context.Context, refreshTokenRequest domain.RefreshTokenRequest) (domain.RefreshTokenResponse, error) {
	collection := r.database.Collection(r.collection)

	var signUpRequest domain.SignUpRequest
	err := collection.FindOne(c, bson.M{"refresh_token": refreshTokenRequest.RefreshToken}).Decode(&signUpRequest)
	if err != nil {
		return domain.RefreshTokenResponse{}, err
	}

	tokenString, err := r.jwtService.GenerateToken(signUpRequest.Username, signUpRequest.Role)
	if err != nil {
		return domain.RefreshTokenResponse{}, err
	}

	return domain.RefreshTokenResponse{
		Message: "Refresh token successful",
		Token:   tokenString,
	}, nil
}
