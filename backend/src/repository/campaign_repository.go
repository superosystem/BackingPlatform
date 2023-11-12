package repository

import (
	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll() ([]entity.Campaign, error)
	FindByUserID(userID int) ([]entity.Campaign, error)
	FindByID(ID int) (entity.Campaign, error)
	Save(campaign entity.Campaign) (entity.Campaign, error)
	Update(campaign entity.Campaign) (entity.Campaign, error)
	CreateImage(campaignImage entity.CampaignImage) (entity.CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)
}

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *campaignRepository {
	return &campaignRepository{db}
}

func (r *campaignRepository) FindAll() ([]entity.Campaign, error) {
	var campaigns []entity.Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *campaignRepository) FindByID(ID int) (entity.Campaign, error) {
	var campaign entity.Campaign

	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) FindByUserID(userID int) ([]entity.Campaign, error) {
	var campaigns []entity.Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *campaignRepository) Save(campaign entity.Campaign) (entity.Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) Update(campaign entity.Campaign) (entity.Campaign, error) {
	err := r.db.Save(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) CreateImage(campaignImage entity.CampaignImage) (entity.CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func (r *campaignRepository) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	err := r.db.Model(&entity.CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
