package service

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"github.com/superosystem/BackingPlatform/backend/src/model"
	"github.com/superosystem/BackingPlatform/backend/src/repository"
)

type CampaignService interface {
	GetCampaigns(userID int) ([]entity.Campaign, error)
	GetCampaignByID(request model.GetCampaignDetailRequest) (entity.Campaign, error)
	CreateCampaign(request model.CreateCampaignRequest) (entity.Campaign, error)
	UpdateCampaign(id model.GetCampaignDetailRequest, request model.CreateCampaignRequest) (entity.Campaign, error)
	SaveCampaignImage(request model.CreateCampaignImageRequest, fileLocation string) (entity.CampaignImage, error)
}

type campaignService struct {
	campaignRepository repository.CampaignRepository
}

func NewCampaignService(campaignRepository repository.CampaignRepository) *campaignService {
	return &campaignService{campaignRepository}
}

func (s *campaignService) GetCampaigns(userID int) ([]entity.Campaign, error) {
	if userID != 0 {
		campaigns, err := s.campaignRepository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.campaignRepository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *campaignService) GetCampaignByID(input model.GetCampaignDetailRequest) (entity.Campaign, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *campaignService) CreateCampaign(input model.CreateCampaignRequest) (entity.Campaign, error) {
	campaign := entity.Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.campaignRepository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *campaignService) UpdateCampaign(inputID model.GetCampaignDetailRequest, inputData model.CreateCampaignRequest) (entity.Campaign, error) {
	campaign, err := s.campaignRepository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updatedCampaign, err := s.campaignRepository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *campaignService) SaveCampaignImage(input model.CreateCampaignImageRequest, fileLocation string) (entity.CampaignImage, error) {
	campaign, err := s.campaignRepository.FindByID(input.CampaignID)
	if err != nil {
		return entity.CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID {
		return entity.CampaignImage{}, errors.New("not an owner of the campaign")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1

		_, err := s.campaignRepository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return entity.CampaignImage{}, err
		}
	}

	campaignImage := entity.CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.campaignRepository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}
