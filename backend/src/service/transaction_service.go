package service

import (
	"errors"
	"strconv"

	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"github.com/superosystem/BackingPlatform/backend/src/model"
	"github.com/superosystem/BackingPlatform/backend/src/repository"
)

type TransactionService interface {
	GetAllTransactions() ([]entity.Transaction, error)
	GetTransactionsByCampaignID(request model.GetCampaignTransactionsRequest) ([]entity.Transaction, error)
	GetTransactionsByUserID(userID int) ([]entity.Transaction, error)
	CreateTransaction(request model.CreateTransactionRequest) (entity.Transaction, error)
	ProcessPayment(request model.TransactionNotifyRequest) error
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
	campaignRepository    repository.CampaignRepository
	paymentService        PaymentService
}

func NewTransactionService(transactionRepository repository.TransactionRepository, campaignRepository repository.CampaignRepository, paymentService PaymentService) *transactionService {
	return &transactionService{transactionRepository, campaignRepository, paymentService}
}

func (s *transactionService) GetAllTransactions() ([]entity.Transaction, error) {
	transactions, err := s.transactionRepository.FindAll()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) GetTransactionsByCampaignID(input model.GetCampaignTransactionsRequest) ([]entity.Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []entity.Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []entity.Transaction{}, errors.New("not an owner of the campaign")
	}

	transactions, err := s.transactionRepository.FindByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) GetTransactionsByUserID(userID int) ([]entity.Transaction, error) {
	transactions, err := s.transactionRepository.FindByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) CreateTransaction(input model.CreateTransactionRequest) (entity.Transaction, error) {
	transaction := entity.Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.UserID = input.User.ID
	transaction.Status = "pending"

	newTransaction, err := s.transactionRepository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := model.Payment{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.transactionRepository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

// for process payment
func (s *transactionService) ProcessPayment(input model.TransactionNotifyRequest) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.transactionRepository.FindByID(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.transactionRepository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindByID(updatedTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}
