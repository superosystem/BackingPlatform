package repository

import (
	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll() ([]entity.Transaction, error)
	FindByCampaignID(campaignID int) ([]entity.Transaction, error)
	FindByUserID(userID int) ([]entity.Transaction, error)
	FindByID(ID int) (entity.Transaction, error)
	Save(transaction entity.Transaction) (entity.Transaction, error)
	Update(transaction entity.Transaction) (entity.Transaction, error)
}

type repo struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) FindAll() ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	err := r.db.Preload("Campaign").Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repo) FindByID(ID int) (entity.Transaction, error) {
	var transaction entity.Transaction

	err := r.db.Where("id = ?", ID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repo) FindByCampaignID(campaignID int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repo) FindByUserID(userID int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repo) Save(transaction entity.Transaction) (entity.Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repo) Update(transaction entity.Transaction) (entity.Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
