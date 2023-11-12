package dto

import (
	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"github.com/superosystem/BackingPlatform/backend/src/model"
)

func CampaignTransactionDTO(transaction entity.Transaction) model.CampaignTransactionResponse {
	mapper := model.CampaignTransactionResponse{}
	mapper.ID = transaction.ID
	mapper.Name = transaction.User.Name
	mapper.Amount = transaction.Amount
	mapper.CreatedAt = transaction.CreatedAt
	return mapper
}

func CampaignTransactionsDTO(transactions []entity.Transaction) []model.CampaignTransactionResponse {
	if len(transactions) == 0 {
		return []model.CampaignTransactionResponse{}
	}

	var transactionsMapper []model.CampaignTransactionResponse
	for _, transaction := range transactions {
		campaginTransactionMapper := CampaignTransactionDTO(transaction)
		transactionsMapper = append(transactionsMapper, campaginTransactionMapper)
	}

	return transactionsMapper
}

func FormatUserTransaction(transaction entity.Transaction) model.UserTransactionResponse {
	mapper := model.UserTransactionResponse{}
	mapper.ID = transaction.ID
	mapper.Amount = transaction.Amount
	mapper.Status = transaction.Status
	mapper.CreatedAt = transaction.CreatedAt

	campaginMapper := model.CampaignUserResponse{}
	campaginMapper.Name = transaction.Campaign.Name
	campaginMapper.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaginMapper.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}
	mapper.Campaign = campaginMapper

	return mapper
}

func UserTransactionsDTO(transactions []entity.Transaction) []model.UserTransactionResponse {
	if len(transactions) == 0 {
		return []model.UserTransactionResponse{}
	}

	var transactionsMapper []model.UserTransactionResponse
	for _, transaction := range transactions {
		userTransactionMapper := FormatUserTransaction(transaction)
		transactionsMapper = append(transactionsMapper, userTransactionMapper)
	}

	return transactionsMapper
}

func TransactionDTO(transaction entity.Transaction) model.TransactionResponse {
	mapper := model.TransactionResponse{}
	mapper.ID = transaction.ID
	mapper.CampaignID = transaction.CampaignID
	mapper.UserID = transaction.UserID
	mapper.Amount = transaction.Amount
	mapper.Status = transaction.Status
	mapper.Code = transaction.Code
	mapper.PaymentURL = transaction.PaymentURL
	return mapper
}
