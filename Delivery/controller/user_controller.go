package controller


import (
	"go_clean_architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	SignUpUsecase domain.SignUpUsecase
}

type LoginController struct {
	LoginUsecase domain.LoginUsecase
}


func (uc *SignUpController) SignUp(c *gin.Context) {
	var signUpRequest domain.SignUpRequest
	if err := c.ShouldBindJSON(&signUpRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	signUpResponse, err := uc.SignUpUsecase.SignUp(c, signUpRequest)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, signUpResponse)
}

func (uc *LoginController) Login(c *gin.Context) {
	var loginRequest domain.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse, err := uc.LoginUsecase.Login(c, loginRequest)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, loginResponse)	
}

// refresh token controller

type RefreshTokenController struct {
	RefreshTokenUsecase domain.RefreshTokenUsecase
}

func (uc *RefreshTokenController) RefreshToken(c *gin.Context) {
	var refreshTokenRequest domain.RefreshTokenRequest
	if err := c.ShouldBindJSON(&refreshTokenRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse, err := uc.RefreshTokenUsecase.RefreshToken(c, refreshTokenRequest)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, refreshTokenResponse)
}




