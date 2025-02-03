package service

import (
	"socmed/dto"
	"socmed/entity"
	"socmed/errorhandle"
	"socmed/repository"
	"time"
)

type ProfileService interface {
	Create(req *dto.ProfileRequest) error
	GetByUserID(userID int) (*dto.ProfileResponse, error)
	Update(profileID int, req *dto.ProfileRequest) error
	Delete(profileID int) error
}

type profileService struct {
	repository repository.ProfileRepository
}

func NewProfileService(r repository.ProfileRepository) *profileService {
	return &profileService{
		repository: r,
	}
}

func (s *profileService) Create(req *dto.ProfileRequest) error {
	profile := entity.Profile{
		UserID:      req.UserID,
		Name:        req.Name,
		Age:         req.Age,
		PhoneNumber: req.PhoneNumber,
	}

	// Save avatar URL if present
	if req.Avatar != nil {
		profile.AvatarUrl = &req.Avatar.Filename
	}

	if err := s.repository.Create(&profile); err != nil {
		return &errorhandle.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *profileService) GetByUserID(userID int) (*dto.ProfileResponse, error) {
	profile, err := s.repository.FindByUserID(userID)
	if err != nil {
		return nil, &errorhandle.InternalServerError{Message: err.Error()}
	}

	// Map the entity to DTO
	profileResponse := &dto.ProfileResponse{
		ID:          profile.ID,
		UserID:      profile.UserID,
		Name:        profile.Name,
		Age:         profile.Age,
		PhoneNumber: profile.PhoneNumber,
		AvatarUrl:   *profile.AvatarUrl,
		CreatedAt:   profile.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   profile.UpdatedAt.Format(time.RFC3339),
	}

	return profileResponse, nil
}

func (s *profileService) Update(profileID int, req *dto.ProfileRequest) error {
	// Find the existing profile
	profile, err := s.repository.FindByID(profileID)
	if err != nil {
		return err
	}

	// Update the profile fields
	profile.Name = req.Name
	profile.Age = req.Age
	profile.PhoneNumber = req.PhoneNumber
	profile.AvatarUrl = &req.AvatarUrl

	// Save the updated profile
	return s.repository.Update(profile)
}

func (s *profileService) Delete(profileID int) error {
	// Check if the profile exists
	_, err := s.repository.FindByID(profileID)
	if err != nil {
		return err
	}

	// Delete the profile
	return s.repository.Delete(profileID)
}
