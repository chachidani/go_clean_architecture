package domain

import "context"
const (
	CollectionUser = "users"
)

// signIn

type SignUpRequest struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Role     string `bson:"role"`
}

type SignUpResponse struct {
	Message string `json:"message"`
}

type SignUpRepository interface {
	SignUp(c context.Context, signUpRequest SignUpRequest) (SignUpResponse, error)
	GetUser(c context.Context) ([]SignUpRequest, error)
}

type SignUpUsecase interface {
	SignUp(c context.Context, signUpRequest SignUpRequest) (SignUpResponse, error)
	GetUser(c context.Context) ([]SignUpRequest, error)
}

// login

type LoginRequest struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type LoginRepository interface {
	Login(c context.Context, loginRequest LoginRequest) (LoginResponse, error)
}

type LoginUsecase interface {
	Login(c context.Context, loginRequest LoginRequest) (LoginResponse, error)
}


// refresh token

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type RefreshTokenRepository interface {
	RefreshToken(c context.Context, refreshTokenRequest RefreshTokenRequest) (RefreshTokenResponse, error)
}

type RefreshTokenUsecase interface {
	RefreshToken(c context.Context, refreshTokenRequest RefreshTokenRequest) (RefreshTokenResponse, error)
}	


