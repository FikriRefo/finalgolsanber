package service

import (
	"socmed/dto"
	"socmed/entity"
	"socmed/errorhandle"
	"socmed/helper"
	"socmed/repository"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (dto.LoginResponse, error) // Perbaikan sintaks
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	if emailExists := s.repository.EmailExists(req.Email); emailExists {
		return &errorhandle.BadRequestError{Message: "Email already exists"}
	}

	if req.Password != req.PasswordConfirm {
		return &errorhandle.BadRequestError{Message: "Password does not match"}
	}

	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandle.InternalServerError{Message: err.Error()}
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: passwordHash,
		Gender:   req.Gender,
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorhandle.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return data, &errorhandle.NotFoundError{Message: "Email atau Password Salah"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return data, &errorhandle.NotFoundError{Message: "Email atau Password Salah"}
	}

	token, err := helper.GenerateToken(user.ID)
	if err != nil {
		return data, &errorhandle.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}

	return data, nil
}
