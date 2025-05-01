package repository

import (
	"context"
	"fmt"
	"go_clean_architecture/domain"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("your-secret-key-here")

// signup repository

type signUpRepository struct {
	database   mongo.Database
	collection string
}

func NewSignUpRepository(database mongo.Database, collection string) domain.SignUpRepository {
	return &signUpRepository{
		database:   database,
		collection: collection,
	}
}

// SignUp implements domain.SignUpRepository.
func (s *signUpRepository) SignUp(c context.Context, signUpRequest domain.SignUpRequest) (domain.SignUpResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.SignUpResponse{}, err
	}

	signUpRequest.Password = string(hashedPassword)

	collection := s.database.Collection(s.collection)
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
	database   mongo.Database
	collection string
}

func NewLoginRepository(database mongo.Database, collection string) domain.LoginRepository {
	return &loginRepository{
		database:   database,
		collection: collection,
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

	if err := bcrypt.CompareHashAndPassword([]byte(signUpRequest.Password), []byte(loginRequest.Password)); err != nil {
		return domain.LoginResponse{}, err
	}

	// Create token with claims
	claims := jwt.MapClaims{
		"username": signUpRequest.Username,
		"role":     signUpRequest.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Log the claims before signing
	fmt.Printf("Generating token with claims: %+v\n", claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Printf("Error signing token: %v\n", err)
		return domain.LoginResponse{}, err
	}

	// Log the generated token
	fmt.Printf("Generated token: %s\n", tokenString)

	return domain.LoginResponse{
		Message: "Login successful as " + signUpRequest.Role,
		Token:   tokenString,
	}, nil
}

// refresh token repository

type refreshTokenRepository struct {
	database   mongo.Database
	collection string
}

func NewRefreshTokenRepository(database mongo.Database, collection string) domain.RefreshTokenRepository {
	return &refreshTokenRepository{
		database:   database,
		collection: collection,
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": signUpRequest.Username,
		"role":     signUpRequest.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return domain.RefreshTokenResponse{}, err
	}

	return domain.RefreshTokenResponse{
		Message: "Refresh token successful",
		Token:   tokenString,
	}, nil
}
