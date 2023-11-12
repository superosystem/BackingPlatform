package dto

import (
	"strings"

	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"github.com/superosystem/BackingPlatform/backend/src/model"
)

func CampaignDTO(campaign entity.Campaign) model.CampaignResponse {
	mapper := model.CampaignResponse{}
	mapper.ID = campaign.ID
	mapper.UserID = campaign.UserID
	mapper.Name = campaign.Name
	mapper.ShortDescription = campaign.ShortDescription
	mapper.GoalAmount = campaign.GoalAmount
	mapper.CurrentAmount = campaign.CurrentAmount
	mapper.Slug = campaign.Slug
	mapper.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		mapper.ImageURL = campaign.CampaignImages[0].FileName
	}

	return mapper
}

func CampaignsDTO(campaigns []entity.Campaign) []model.CampaignResponse {
	campaignsMapper := []model.CampaignResponse{}

	for _, campaign := range campaigns {
		campaignMapper := CampaignDTO(campaign)
		campaignsMapper = append(campaignsMapper, campaignMapper)
	}

	return campaignsMapper
}

func CampaignDetailDTO(campaign entity.Campaign) model.CampaignDetailResponse {
	mapper := model.CampaignDetailResponse{}
	mapper.ID = campaign.ID
	mapper.Name = campaign.Name
	mapper.ShortDescription = campaign.ShortDescription
	mapper.Description = campaign.Description
	mapper.GoalAmount = campaign.GoalAmount
	mapper.CurrentAmount = campaign.CurrentAmount
	mapper.BackerCount = campaign.BackerCount
	mapper.UserID = campaign.UserID
	mapper.Slug = campaign.Slug
	mapper.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		mapper.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	mapper.Perks = perks

	user := campaign.User

	campaignUserMapper := model.CampaignUserResponse{}
	campaignUserMapper.Name = user.Name
	campaignUserMapper.ImageURL = user.AvatarFileName

	mapper.User = campaignUserMapper

	images := []model.CampaignImageResponse{}
	for _, image := range campaign.CampaignImages {
		campaignImageMapper := model.CampaignImageResponse{}
		campaignImageMapper.ImageURL = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageMapper.IsPrimary = isPrimary
		images = append(images, campaignImageMapper)
	}
	mapper.Images = images

	return mapper
}
