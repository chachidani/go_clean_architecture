package usecases

import (
	"context"
	"go_clean_architecture/domain"
	"time"
)

// signupusecase

type signUpUsecase struct {
	signUpRepository domain.SignUpRepository
	ContextTimeout   time.Duration
}

func NewSignUpUsecase(signUpRepository domain.SignUpRepository, timeout time.Duration) domain.SignUpUsecase {
	return &signUpUsecase{
		signUpRepository: signUpRepository,
		ContextTimeout:   timeout,
	}
}

// SignUp implements domain.SignUpUsecase.
func (s *signUpUsecase) SignUp(c context.Context, signUpRequest domain.SignUpRequest) (domain.SignUpResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.ContextTimeout)
	defer cancel()

	return s.signUpRepository.SignUp(ctx, signUpRequest)
}

func (s *signUpUsecase) GetUser(c context.Context) ([]domain.SignUpRequest, error) {
	ctx, cancel := context.WithTimeout(c, s.ContextTimeout)
	defer cancel()

	return s.signUpRepository.GetUser(ctx)
}


// loginusecase

type loginUsecase struct {
	loginRepository domain.LoginRepository
	ContextTimeout  time.Duration
}


func NewLoginUsecase(loginRepository domain.LoginRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		loginRepository: loginRepository,
		ContextTimeout:  timeout,
	}
}

// Login implements domain.LoginUsecase.
func (l *loginUsecase) Login(c context.Context, loginRequest domain.LoginRequest) (domain.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(c, l.ContextTimeout)
	defer cancel()

	return l.loginRepository.Login(ctx, loginRequest)
}




// refresh token usecase

type refreshTokenUsecase struct {
	refreshTokenRepository domain.RefreshTokenRepository
	ContextTimeout       time.Duration
}

func NewRefreshTokenUsecase(refreshTokenRepository domain.RefreshTokenRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		refreshTokenRepository: refreshTokenRepository,
		ContextTimeout:         timeout,
	}
}	

// RefreshToken implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) RefreshToken(c context.Context, refreshTokenRequest domain.RefreshTokenRequest) (domain.RefreshTokenResponse, error) {
	ctx, cancel := context.WithTimeout(c, r.ContextTimeout)
	defer cancel()	

	return r.refreshTokenRepository.RefreshToken(ctx, refreshTokenRequest)
}


	
