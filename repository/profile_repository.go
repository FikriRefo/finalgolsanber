package repository

import (
	"socmed/entity"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindByID(profileID int) (*entity.Profile, error)
	Create(profile *entity.Profile) error
	Update(profile *entity.Profile) error
	Delete(profileID int) error
	FindByUserID(userID int) (*entity.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *profileRepository {
	return &profileRepository{
		db: db,
	}
}

func (r *profileRepository) FindByID(profileID int) (*entity.Profile, error) {
	var profile entity.Profile
	err := r.db.Where("id = ?", profileID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *profileRepository) Create(profile *entity.Profile) error {
	return r.db.Create(profile).Error
}

func (r *profileRepository) Update(profile *entity.Profile) error {
	return r.db.Save(profile).Error
}

func (r *profileRepository) Delete(profileID int) error {
	err := r.db.Where("id = ?", profileID).Delete(&entity.Profile{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *profileRepository) FindByUserID(userID int) (*entity.Profile, error) {
	var profile entity.Profile
	err := r.db.Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
